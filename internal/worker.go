package internal

import (
	"fmt"
	"time"
)

type Worker struct {
	ID         int      
	Queue      *Queue   
	JobChannel chan Job 
}


func NewWorker(id int, queue *Queue) *Worker {
	return &Worker{
		ID:         id,
		Queue:      queue,
		JobChannel: queue.JobChannel,
	}
}


func (w *Worker) Start() {
	go func() {
		for job := range w.JobChannel {
			
			fmt.Printf("Worker %d started processing job: %s (ID: %s)\n", w.ID, job.Name, job.ID)
			
			select {
			case <-time.After(time.Duration(job.Timeout) * time.Second):
				
				fmt.Printf("Worker %d timed out on job: %s (ID: %s)\n", w.ID, job.Name, job.ID)
			}
			
			fmt.Printf("Worker %d finished processing job: %s (ID: %s)\n", w.ID, job.Name, job.ID)
		}
	}()
}
