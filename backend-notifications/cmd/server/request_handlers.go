package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/none-da/otel-tryouts/backend-notifications/pkg/notifications"
	"go.opentelemetry.io/contrib/instrumentation/net/http/httptrace/otelhttptrace"
	"go.opentelemetry.io/otel/api/baggage"
	"go.opentelemetry.io/otel/api/global"
	"go.opentelemetry.io/otel/api/trace"
)

func handleErrorResponse(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	fmt.Fprintf(w, err.Error())
}

func getNotifications(w http.ResponseWriter, r *http.Request) {
	requestContext := r.Context()
	attrs, entries, spanCtx := otelhttptrace.Extract(requestContext, r)
	if spanCtx.IsValid() {
		requestContext = trace.ContextWithRemoteSpanContext(requestContext, spanCtx)
	}
	r = r.WithContext(baggage.ContextWithMap(requestContext, baggage.NewMap(baggage.MapUpdate{
		MultiKV: entries,
	})))
	tracer := global.Tracer("notifications")
	_, span := tracer.Start(
		r.Context(),
		"/",
		trace.WithAttributes(attrs...),
	)
	defer span.End()

	notificationsJSON, err := json.Marshal(notifications.GetNotifications())
	if err != nil {
		handleErrorResponse(w, errors.New("Error! Notifications can't be retrieved"))
		return
	}
	fmt.Fprintf(w, string(notificationsJSON))
}
