package main

import (
	"fmt"
	"net/http"
	"time"

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
	fn := initTracer(serviceName)
	defer fn()
	log.Info("Running Notifications Service on ", port, "...")
	http.HandleFunc("/", getNotifications)
	http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}
