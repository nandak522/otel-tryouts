package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/none-da/otel-tryouts/backend-tweets/pkg/tweets"
)

func handleErrorResponse(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	fmt.Fprintf(w, err.Error())
}

func getTweets(w http.ResponseWriter, r *http.Request) {
	tweetsJSON, err := json.Marshal(tweets.GetTweets())
	if err != nil {
		handleErrorResponse(w, errors.New("Error! Tweets can't be retrieved"))
		return
	}
	fmt.Fprintf(w, string(tweetsJSON))
}
