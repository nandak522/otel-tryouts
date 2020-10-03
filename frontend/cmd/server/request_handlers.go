package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/newrelic/go-agent/v3/newrelic"
	assembler "github.com/none-da/otel-tryouts/frontend/pkg/assembler"
)

func handleErrorResponse(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	fmt.Fprintf(w, err.Error())
}

func homepage(w http.ResponseWriter, r *http.Request) {
	txn := newrelic.FromContext(r.Context())
	defer txn.End()

	data, err := json.Marshal(assembler.GetData(txn))
	if err != nil {
		handleErrorResponse(w, errors.New("msg couldn't be saved. Reason:"+err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, string(data))
}
