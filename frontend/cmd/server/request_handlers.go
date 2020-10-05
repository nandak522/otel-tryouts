package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	apm "github.com/newrelic/go-agent/v3/newrelic"
	"go.opentelemetry.io/contrib/instrumentation/net/http/httptrace/otelhttptrace"
	"go.opentelemetry.io/otel/api/global"
)

func handleErrorResponse(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	fmt.Fprintf(w, err.Error())
}

func homepage(w http.ResponseWriter, r *http.Request) {
	txn := apm.FromContext(r.Context())
	defer txn.End()

	requestContext := r.Context()
	frontEndClient := http.DefaultClient
	var tweets string
	var notifications string
	tracer := global.Tracer("frontend")
	rootSpanContext, rootSpan := tracer.Start(requestContext, "/homepage")
	req, _ := http.NewRequest("GET", "http://localhost:8001", nil)
	otelhttptrace.Inject(rootSpanContext, req)

	tweetsResponse, err := frontEndClient.Do(req)
	if err != nil {
		fmt.Println("Error from Tweets Service")
	}
	defer tweetsResponse.Body.Close()

	body, err := ioutil.ReadAll(tweetsResponse.Body)
	if err != nil {
		fmt.Println("Error in reading tweetsResponse.Body")
	}
	tweets = string(body)

	req, _ = http.NewRequest("GET", "http://localhost:8002", nil)
	otelhttptrace.Inject(rootSpanContext, req)
	notificationsResponse, err := frontEndClient.Do(req)
	if err != nil {
		fmt.Println("Error from Notifications Service")
	}
	defer notificationsResponse.Body.Close()

	body, err = ioutil.ReadAll(notificationsResponse.Body)
	if err != nil {
		fmt.Println("Error in reading notificationsResponse.Body")
	}
	notifications = string(body)

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
