package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"

	apm "github.com/newrelic/go-agent/v3/newrelic"
	"go.opentelemetry.io/contrib/instrumentation/net/http/httptrace/otelhttptrace"
	"go.opentelemetry.io/otel/api/global"
)

func handleErrorResponse(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	fmt.Fprintf(w, err.Error())
}

func makeExternalCall(rootSpanContext *context.Context, url string, wg *sync.WaitGroup) string {
	defer wg.Done()
	req, _ := http.NewRequest("GET", url, nil)
	otelhttptrace.Inject(*rootSpanContext, req)
	frontEndClient := http.DefaultClient
	tweetsResponse, err := frontEndClient.Do(req)
	if err != nil {
		fmt.Println("Error from Tweets Service")
	}
	defer tweetsResponse.Body.Close()

	body, err := ioutil.ReadAll(tweetsResponse.Body)
	if err != nil {
		fmt.Println("Error in reading tweetsResponse.Body")
	}
	return string(body)
}

func computeSomethingLocal(rootSpanContext *context.Context) {
	tracer := global.Tracer("local-compute")

	_, trace := tracer.Start(*rootSpanContext, "compute-something")
	time.Sleep(50 * time.Millisecond)

	trace.End()
}

func homepage(w http.ResponseWriter, r *http.Request) {
	txn := apm.FromContext(r.Context())
	defer txn.End()

	requestContext := r.Context()
	tracer := global.Tracer("frontend")
	rootSpanContext, rootSpan := tracer.Start(requestContext, "/homepage")

	var tweets string
	var notifications string
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		tweets = makeExternalCall(&rootSpanContext, "http://localhost:8001", &wg)
	}()
	wg.Add(1)
	go func() {
		notifications = makeExternalCall(&rootSpanContext, "http://localhost:8002", &wg)
	}()
	computeSomethingLocal(&rootSpanContext)
	wg.Wait()

	rootSpan.End()

	homepageData := make(map[string]string)
	homepageData["tweets"] = tweets
	homepageData["notifications"] = notifications
	data, err := json.Marshal(homepageData)
	if err != nil {
		handleErrorResponse(w, errors.New("msg couldn't be saved. Reason:"+err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, string(data))
}
