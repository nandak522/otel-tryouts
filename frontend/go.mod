module github.com/none-da/otel-tryouts/frontend

go 1.15

require (
	github.com/newrelic/go-agent/v3 v3.9.0
	github.com/sirupsen/logrus v1.7.0
	github.com/spf13/pflag v1.0.5
	go.opentelemetry.io/contrib/instrumentation/net/http/httptrace/otelhttptrace v0.12.0
	go.opentelemetry.io/otel v0.12.0
	go.opentelemetry.io/otel/exporters/trace/jaeger v0.12.0
	go.opentelemetry.io/otel/sdk v0.12.0
)
