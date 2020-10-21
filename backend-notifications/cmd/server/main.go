package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/newrelic/opentelemetry-exporter-go/newrelic"
	log "github.com/sirupsen/logrus"
	flag "github.com/spf13/pflag"
)

func main() {
	defaultPort := "8002"
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
	serviceName := "notifications"
	// jaegerFn := initTracer(serviceName)
	// defer jaegerFn()
	// Assumes the NEW_RELIC_API_KEY environment variable contains your New
	// Relic Insights insert API key. This will error if it does not.
	// newrelicFn := initNewrelicTracer(serviceName)
	// defer newrelicFn.Stop()
	controller, err := newrelic.InstallNewPipeline(serviceName)
	if err != nil {
		panic(err)
	}
	defer controller.Stop()

	log.Info("Running Notifications Service on ", port, "...")
	http.HandleFunc("/", getNotifications)
	http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}
