package assembler

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetTweets() string {
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

func GetNotifications() string {
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
// func GetData(requestContext context.Context, txn *apm.Transaction) map[string]string {
// 	data := make(map[string]string)

// 	getTweetsSegment := apm.Segment{}
// 	getTweetsSegment.Name = "getTweets"
// 	getTweetsSegment.StartTime = txn.StartSegmentNow()

// 	requestContext, req = httptrace.W3C(requestContext, req)
// 	httptrace.Inject(requestContext, req)

// 	tracer := global.Tracer("homepage-tracer")
// 	_, getTweetsspan := tracer.Start(requestContext, "/tweets")
// 	data["tweets"] = getTweets()
// 	getTweetsspan.End()

// 	getTweetsSegment.End()

// 	getNotificationsSegment := apm.Segment{}
// 	getNotificationsSegment.Name = "getNotifications"
// 	getNotificationsSegment.StartTime = txn.StartSegmentNow()

// 	_, getNotificationsspan := tracer.Start(requestContext, "/notifications")
// 	data["notifications"] = getNotifications()
// 	getNotificationsspan.End()

// 	getNotificationsSegment.End()
// 	return data
// }
