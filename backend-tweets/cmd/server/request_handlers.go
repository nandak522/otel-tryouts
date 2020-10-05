package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/none-da/otel-tryouts/backend-tweets/pkg/tweets"
	"go.opentelemetry.io/contrib/instrumentation/net/http/httptrace/otelhttptrace"
	"go.opentelemetry.io/otel/api/baggage"
	"go.opentelemetry.io/otel/api/global"
	"go.opentelemetry.io/otel/api/trace"
)

func handleErrorResponse(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	fmt.Fprintf(w, err.Error())
}

func getTweets(w http.ResponseWriter, r *http.Request) {
	requestContext := r.Context()
	attrs, entries, spanCtx := otelhttptrace.Extract(requestContext, r)
	if spanCtx.IsValid() {
		requestContext = trace.ContextWithRemoteSpanContext(requestContext, spanCtx)
	}
	r = r.WithContext(baggage.ContextWithMap(requestContext, baggage.NewMap(baggage.MapUpdate{
		MultiKV: entries,
	})))
	tracer := global.Tracer("tweets")
	_, span := tracer.Start(
		r.Context(),
		"/",
		trace.WithAttributes(attrs...),
	)
	defer span.End()

	tweetsJSON, err := json.Marshal(tweets.GetTweets())
	if err != nil {
		handleErrorResponse(w, errors.New("Error! Tweets can't be retrieved"))
		return
	}
	fmt.Fprintf(w, string(tweetsJSON))
}
