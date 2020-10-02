package notifications

import "time"

// GetNotifications ...
func GetNotifications() []string {
	time.Sleep(250 * time.Millisecond) // Analogous to a db call
	notifications := []string{
		"Read this 1",
		"Read this 2",
		"Read this 3",
		"Read this 4",
	}
	return notifications
}
