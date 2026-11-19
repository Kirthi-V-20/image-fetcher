package downloader

import (
	"fmt"
	"net/http"

	"image-fetcher/internal/storage"
)

func DownloadImage(imageID int, imageURL string) error {

	response, err := http.Get(imageURL)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("download failed: %s", response.Status)
	}

	err = storage.SaveImage(imageID, response.Body)
	if err != nil {
		return err
	}

	fmt.Printf("Image %d downloaded successfully\n", imageID)

	return nil
}
