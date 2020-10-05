package main

import (
	"os"

	log "github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel/exporters/trace/jaeger"

	"go.opentelemetry.io/otel/label"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

func initTracer(service string) func() {
	log.Info("Will post traces to ", os.Getenv("JAEGER_COLLECTOR"))
	flush, err := jaeger.InstallNewPipeline(
		jaeger.WithCollectorEndpoint(os.Getenv("JAEGER_COLLECTOR")),
		jaeger.WithProcess(jaeger.Process{
			ServiceName: service,
			Tags: []label.KeyValue{
				label.Key("exporter").String("jaeger"),
			},
		}),
		jaeger.WithSDK(&sdktrace.Config{DefaultSampler: sdktrace.AlwaysSample()}),
	)
	if err != nil {
		log.Fatal(err)
	}

	return func() {
		flush()
		log.Info("initTracer Flush happened")
	}
}
