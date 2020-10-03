package assembler

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/newrelic/go-agent/v3/newrelic"
)

func getTweets() string {
	tweetsResponse, err := http.Get("http://localhost:8001")
	if err != nil {
		fmt.Println("Error from Tweets Service")
	}
	defer tweetsResponse.Body.Close()

	body, err := ioutil.ReadAll(tweetsResponse.Body)
	if err != nil {
		fmt.Println("Error in reading tweetsResponse.Body")
	}
	return string(body)
}

func getNotifications() string {
	notificationsResponse, err := http.Get("http://localhost:8002")
	if err != nil {
		fmt.Println("Error from Notifications Service")
	}
	defer notificationsResponse.Body.Close()

	body, err := ioutil.ReadAll(notificationsResponse.Body)
	if err != nil {
		fmt.Println("Error in reading notificationsResponse.Body")
	}
	return string(body)
}

// GetData ...
func GetData(txn *newrelic.Transaction) map[string]string {
	data := make(map[string]string)

	getTweetsSegment := newrelic.Segment{}
	getTweetsSegment.Name = "getTweets"
	getTweetsSegment.StartTime = txn.StartSegmentNow()
	data["tweets"] = getTweets()
	getTweetsSegment.End()

	getNotificationsSegment := newrelic.Segment{}
	getNotificationsSegment.Name = "getNotifications"
	getNotificationsSegment.StartTime = txn.StartSegmentNow()
	data["notifications"] = getNotifications()
	getNotificationsSegment.End()
	return data
}
