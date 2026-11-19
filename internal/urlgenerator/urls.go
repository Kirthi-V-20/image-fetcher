package urlgenerator

import "fmt"

func GenerateURLs(totalImages int) []string {
	var urls []string

	for i := 1; i <= totalImages; i++ {
		url := fmt.Sprintf(
			"https://picsum.photos/seed/%d/400/300",
			i,
		)

		urls = append(urls, url)
	}

	return urls
}
