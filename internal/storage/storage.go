package storage

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func SaveImage(imageID int, imageData io.Reader) error {

	err := os.MkdirAll("images", 0755)
	if err != nil {
		return err
	}

	fileName := fmt.Sprintf("image_%03d.jpg", imageID)
	filePath := filepath.Join("images", fileName)

	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, imageData)
	if err != nil {
		return err
	}

	return nil
}
