package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	log "github.com/sirupsen/logrus"
	flag "github.com/spf13/pflag"

	"github.com/newrelic/go-agent/v3/newrelic"
	// otel_newrelic "github.com/newrelic/opentelemetry-exporter-go/newrelic"
)

// StartTime gives the start time of server
// var StartTime = time.Now()

// func uptime() string {
// 	elapsedTime := time.Since(StartTime)
// 	return fmt.Sprintf("%d:%d:%d", int(math.Round(elapsedTime.Hours())), int(math.Round(elapsedTime.Minutes())), int(math.Round(elapsedTime.Seconds())))
// }

// func initTracer() {
// 	exporter, err := otel_newrelic.NewExporter("twitter-frontend", os.Getenv("NEWRELIC_LICENSE_KEY"))
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	tp := trace.NewTracerProvider(trace.WithSyncer(exporter))
// 	global.SetTracerProvider(tp)
// }

func main() {
	defaultPort := "8000"
	var port string
	flag.StringVarP(&port, "port", "p", defaultPort, "Port. Defaults to "+defaultPort)
	var printHelp bool
	flag.BoolVarP(&printHelp, "help", "h", false, "Prints this help content.")
	flag.Parse()
	if printHelp {
		flag.Usage()
		return
	}
	log.SetFormatter(&log.TextFormatter{
		DisableLevelTruncation: true,
		PadLevelText:           true,
		TimestampFormat:        time.RFC3339,
		FullTimestamp:          true,
		ForceColors:            false,
	})

	log.SetLevel(log.DebugLevel)
	fn := initTracer()
	defer fn()
	log.Info("Running Frontend Service on ", port, "...")
	newrelicLicenseKey, isEnvVarSet := os.LookupEnv("NEWRELIC_LICENSE_KEY")
	if isEnvVarSet {
		var newrelicAPM string
		newrelicAPM, isEnvVarSet := os.LookupEnv("NEWRELIC_APM")
		if !isEnvVarSet {
			newrelicAPM = "twitter-frontend"
		}
		app, _ := newrelic.NewApplication(
			newrelic.ConfigAppName(newrelicAPM),
			newrelic.ConfigLicense(newrelicLicenseKey),
			// newrelic.ConfigDebugLogger(os.Stdout), // Only when debugging a problem
		)
		log.Info("Newrelic Monitoring is enabled. Posting to ", newrelicAPM, " APM")
		http.HandleFunc(newrelic.WrapHandleFunc(app, "/", homepage))
	} else {
		log.Info("NEWRELIC_LICENSE_KEY env variable not found. Instrumentation is off. Moving on...")
		http.HandleFunc("/", homepage)
	}
	http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}
