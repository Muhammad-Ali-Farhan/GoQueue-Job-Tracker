package storage

import (
	"os"
	"encoding/json"
	"fmt"
	"goqueue/internal"
)

func SaveJobToFile(job *internal.Job) error {
	file, err := os.OpenFile("jobs.json", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	jobData, err := json.Marshal(job)
	if err != nil {
		return err
	}

	if _, err := file.Write(append(jobData, '\n')); err != nil {
		return err
	}

	fmt.Printf("Job saved to file: %s\n", job.ID)
	return nil
}
