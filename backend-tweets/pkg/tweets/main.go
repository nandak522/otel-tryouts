package tweets

import "time"

// GetTweets ...
func GetTweets() []string {
	time.Sleep(250 * time.Millisecond) // Analogous to a db call
	tweets := []string{
		"Tweet 1",
		"Tweet 2",
		"Tweet 3",
		"Tweet 4",
		"Tweet 5",
	}
	return tweets
}
