package main

import (
	"github.com/newrelic/opentelemetry-exporter-go/newrelic"
	log "github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel/sdk/metric/controller/push"
)

func initNewrelicTracer(service string) *push.Controller {
	// Assumes the NEW_RELIC_API_KEY environment variable contains your New
	// Relic Insights insert API key. This will error if it does not.
	controller, err := newrelic.InstallNewPipeline(service)
	if err != nil {
		log.Fatal(err)
	}
	return controller
}
