package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/none-da/otel-tryouts/backend-tweets/pkg/tweets"
	"go.opentelemetry.io/otel/api/global"
)

func handleErrorResponse(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	fmt.Fprintf(w, err.Error())
}

func getTweets(w http.ResponseWriter, r *http.Request) {
	tracer := global.Tracer("notifications")
	requestContext := r.Context()
	_, rootSpan := tracer.Start(requestContext, "/notifications")

	tweetsJSON, err := json.Marshal(tweets.GetTweets())
	if err != nil {
		handleErrorResponse(w, errors.New("Error! Tweets can't be retrieved"))
		return
	}
	fmt.Fprintf(w, string(tweetsJSON))
	rootSpan.End()
}
