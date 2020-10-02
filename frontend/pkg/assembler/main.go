package assembler

import (
	"fmt"
	"io/ioutil"
	"net/http"
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
func GetData() map[string]string {
	data := make(map[string]string)
	data["tweets"] = getTweets()
	data["notifications"] = getNotifications()
	return data
}
