package main

import (
	"os"

	log "github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel/exporters/trace/jaeger"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

// initTracer creates a new trace provider instance and registers it as global trace provider.
func initTracer(serviceName string) func() {
	flush, err := jaeger.InstallNewPipeline(
		jaeger.WithCollectorEndpoint(os.Getenv("JAEGER_COLLECTOR")),
		jaeger.WithProcess(jaeger.Process{
			ServiceName: serviceName,
			// Tags: []label.KeyValue{
			// 	label.String("exporter", "jaeger"),
			// 	label.Float64("float", 312.23),
			// },
		}),
		jaeger.WithSDK(&sdktrace.Config{DefaultSampler: sdktrace.AlwaysSample()}),
	)
	if err != nil {
		log.Fatal(err)
	}

	return func() {
		flush()
	}
}
