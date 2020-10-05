package main

import (
	"os"

	log "github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel/exporters/trace/jaeger"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

// func initExporter() {
// 	exporter, _ := otlp.NewExporter()
// 	defer exporter.Shutdown(context.Background())

// 	tp := sdktrace.NewTracerProvider(
// 		sdktrace.WithConfig(sdktrace.Config{DefaultSampler: sdktrace.AlwaysSample()}),
// 		sdktrace.WithBatcher(
// 			exporter,
// 			// add following two options to ensure flush
// 			sdktrace.WithBatchTimeout(5),
// 			sdktrace.WithMaxExportBatchSize(10),
// 		),
// 	)
// 	global.SetTracerProvider(tp)
// }

// func initExporter() {
// 	exporter, err := otel_newrelic.NewExporter("twitter-frontend", os.Getenv("NEWRELIC_LICENSE_KEY"))
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	tp := trace.NewTracerProvider(trace.WithSyncer(exporter))
// 	global.SetTracerProvider(tp)
// }

// initTracer creates a new trace provider instance and registers it as global trace provider.
func initTracer() func() {

	// Create and install Jaeger export pipeline
	flush, err := jaeger.InstallNewPipeline(
		jaeger.WithCollectorEndpoint(os.Getenv("JAEGER_COLLECTOR")),
		jaeger.WithProcess(jaeger.Process{
			ServiceName: "twitter-frontend",
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
