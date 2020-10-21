package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/newrelic/opentelemetry-exporter-go/newrelic"
	log "github.com/sirupsen/logrus"
	flag "github.com/spf13/pflag"
)

// StartTime gives the start time of server
// var StartTime = time.Now()

// func uptime() string {
// 	elapsedTime := time.Since(StartTime)
// 	return fmt.Sprintf("%d:%d:%d", int(math.Round(elapsedTime.Hours())), int(math.Round(elapsedTime.Minutes())), int(math.Round(elapsedTime.Seconds())))
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
	serviceName := "frontend"

	// jaegerFn := initTracer(serviceName)
	// defer jaegerFn()

	// Assumes the NEW_RELIC_API_KEY environment variable contains your New
	// Relic Insights insert API key. This will error if it does not.
	controller, err := newrelic.InstallNewPipeline(serviceName)
	if err != nil {
		panic(err)
	}
	defer controller.Stop()

	log.Info("Running Frontend Service on ", port, "...")
	// apmName := "twitter-frontend"
	// path, handler := initAPM(apmName)
	// http.HandleFunc(path, handler)
	http.HandleFunc("/", homepage)
	http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}
