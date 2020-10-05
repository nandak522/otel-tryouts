package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/none-da/otel-tryouts/backend-notifications/pkg/notifications"
	"go.opentelemetry.io/otel/api/global"
)

func handleErrorResponse(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	fmt.Fprintf(w, err.Error())
}

func getNotifications(w http.ResponseWriter, r *http.Request) {
	tracer := global.Tracer("notifications")
	requestContext := r.Context()
	_, rootSpan := tracer.Start(requestContext, "/notifications")
	notificationsJSON, err := json.Marshal(notifications.GetNotifications())
	if err != nil {
		handleErrorResponse(w, errors.New("Error! Notifications can't be retrieved"))
		return
	}
	fmt.Fprintf(w, string(notificationsJSON))
	rootSpan.End()
}
