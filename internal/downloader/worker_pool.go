package downloader

import (
	"fmt"
	"sync"
)

type DownloadJob struct {
	ID  int
	URL string
}

func StartWorkerPool(urls []string, workerCount int) error {

	if workerCount <= 0 {
		return fmt.Errorf("worker count must be greater than zero")
	}

	jobs := make(chan DownloadJob)

	var wg sync.WaitGroup

	wg.Add(workerCount)

	for i := 1; i <= workerCount; i++ {
		go func() {
			defer wg.Done()

			for job := range jobs {

				err := DownloadImage(job.ID, job.URL)
				if err != nil {
					fmt.Printf("Image %d failed: %v\n", job.ID, err)
				}
			}
		}()
	}

	for i, url := range urls {
		jobs <- DownloadJob{
			ID:  i + 1,
			URL: url,
		}
	}

	close(jobs)

	wg.Wait()

	return nil
}
