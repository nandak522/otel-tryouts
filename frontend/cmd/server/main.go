package main

import (
	"fmt"
	"net/http"
	"time"

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
	log.Info("Running Frontend Service on ", port, "...")
	http.HandleFunc("/", homepage)
	http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}
