package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/none-da/otel-tryouts/backend-tweets/pkg/tweets"
	log "github.com/sirupsen/logrus"
	"go.opentelemetry.io/contrib/instrumentation/net/http/httptrace/otelhttptrace"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/api/global"
	"go.opentelemetry.io/otel/api/trace"
	"go.opentelemetry.io/otel/propagators"
)

func handleErrorResponse(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	fmt.Fprintf(w, err.Error())
}

func getTweets(w http.ResponseWriter, r *http.Request) {
	requestContext := r.Context()
	attrs, entries, spanCtx := otelhttptrace.Extract(requestContext, r, otelhttptrace.WithPropagators(otel.NewCompositeTextMapPropagator(propagators.TraceContext{}, propagators.Baggage{})))
	log.Debug("tweets spanCtx.TraceID: ", spanCtx.TraceID)
	if spanCtx.IsValid() {
		requestContext = trace.ContextWithRemoteSpanContext(requestContext, spanCtx)
	}
	r = r.WithContext(otel.ContextWithBaggageValues(requestContext, entries...))
	tracer := global.Tracer("tweets")
	_, span := tracer.Start(
		r.Context(),
		"/",
		trace.WithAttributes(attrs...),
	)
	log.Debug("tweets span.SpanContext().SpanID: ", span.SpanContext().SpanID)
	defer span.End()

	tweetsJSON, err := json.Marshal(tweets.GetTweets())
	if err != nil {
		handleErrorResponse(w, errors.New("Error! Tweets can't be retrieved"))
		return
	}
	fmt.Fprintf(w, string(tweetsJSON))
}
