package main

import (
	"fmt"

	"image-fetcher/internal/downloader"
	"image-fetcher/internal/urlgenerator"
)

func main() {
	urls := urlgenerator.GenerateURLs(100)

	fmt.Println("Starting image downloads...")

	err := downloader.StartWorkerPool(urls, 10)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("All download jobs completed!")
}
