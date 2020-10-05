package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	apm "github.com/newrelic/go-agent/v3/newrelic"
	assembler "github.com/none-da/otel-tryouts/frontend/pkg/assembler"
	"go.opentelemetry.io/otel/api/global"
)

func handleErrorResponse(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	fmt.Fprintf(w, err.Error())
}

func homepage(w http.ResponseWriter, r *http.Request) {
	txn := apm.FromContext(r.Context())
	defer txn.End()

	tracer := global.Tracer("homepage-tracer")
	_, span := tracer.Start(r.Context(), "/homepage")
	data, err := json.Marshal(assembler.GetData(txn))
	span.End()

	if err != nil {
		handleErrorResponse(w, errors.New("msg couldn't be saved. Reason:"+err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, string(data))
}
