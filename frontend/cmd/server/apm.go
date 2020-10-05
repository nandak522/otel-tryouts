package main

import (
	"net/http"
	"os"

	apm "github.com/newrelic/go-agent/v3/newrelic"
	log "github.com/sirupsen/logrus"
)

func initAPM(serviceName string) (string, func(http.ResponseWriter, *http.Request)) {
	newrelicLicenseKey, isEnvVarSet := os.LookupEnv("NEWRELIC_LICENSE_KEY")
	if isEnvVarSet {
		var newrelicAPM string
		newrelicAPM, isEnvVarSet := os.LookupEnv("NEWRELIC_APM")
		if !isEnvVarSet {
			newrelicAPM = serviceName
		}
		app, _ := apm.NewApplication(
			apm.ConfigAppName(newrelicAPM),
			apm.ConfigLicense(newrelicLicenseKey),
			// apm.ConfigDebugLogger(os.Stdout), // Only when debugging a problem
		)
		log.Info("Newrelic Monitoring is enabled. Posting to ", newrelicAPM, " APM")
		return apm.WrapHandleFunc(app, "/", homepage)
	}
	log.Info("NEWRELIC_LICENSE_KEY env variable not found. APM Instrumentation is off. Moving on...")
	return "/", homepage
}
