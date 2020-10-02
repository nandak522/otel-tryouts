package assembler

// GetData ...
func GetData() map[string][]string {
	tweets := []string{
		"Tweet 1",
		"Tweet 2",
		"Tweet 3",
		"Tweet 4",
		"Tweet 5",
	}
	notifications := []string{
		"Read this 1",
		"Read this 2",
		"Read this 3",
		"Read this 4",
	}
	data := make(map[string][]string)
	data["tweets"] = tweets
	data["notifications"] = notifications
	return data
}
