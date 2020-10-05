module github.com/none-da/otel-tryouts/frontend

go 1.15

require (
	github.com/newrelic/go-agent/v3 v3.9.0
	github.com/newrelic/opentelemetry-exporter-go v0.1.1-0.20201002153017-2c934cc28388
	github.com/sirupsen/logrus v1.7.0
	github.com/spf13/pflag v1.0.5
	go.opentelemetry.io/otel v0.12.0
	go.opentelemetry.io/otel/exporters/otlp v0.12.0
	go.opentelemetry.io/otel/exporters/trace/jaeger v0.12.0
	go.opentelemetry.io/otel/sdk v0.12.0
)
