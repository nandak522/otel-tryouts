package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	assembler "github.com/none-da/otel-tryouts/frontend/pkg/assembler"
)

func handleErrorResponse(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	fmt.Fprintf(w, err.Error())
}

func homepage(w http.ResponseWriter, r *http.Request) {
	data, err := json.Marshal(assembler.GetData())
	if err != nil {
		handleErrorResponse(w, errors.New("msg couldn't be saved. Reason:"+err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, string(data))
}
