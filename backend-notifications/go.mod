module github.com/none-da/otel-tryouts/backend-notifications

go 1.15

require (
	github.com/newrelic/opentelemetry-exporter-go v0.1.1-0.20201015231732-c523eeb166d5
	github.com/sirupsen/logrus v1.7.0
	github.com/spf13/pflag v1.0.5
	go.opentelemetry.io/contrib/instrumentation/net/http/httptrace/otelhttptrace v0.13.0
	go.opentelemetry.io/otel v0.13.0
	go.opentelemetry.io/otel/exporters/trace/jaeger v0.13.0
	go.opentelemetry.io/otel/sdk v0.13.0
)
