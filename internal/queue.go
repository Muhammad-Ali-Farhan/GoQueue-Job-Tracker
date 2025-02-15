package internal

import (
	"fmt"
	"sync"
)


type Job struct {
	ID      string 
	Name    string 
	Timeout int    
}


type Queue struct {
	Jobs        []Job     
	JobChannel  chan Job  
	WorkerPool  int       
	mu          sync.Mutex 
	Workers     []*Worker 
}


func NewQueue(workerPool int) *Queue {
	return &Queue{
		Jobs:       []Job{},
		JobChannel: make(chan Job),
		WorkerPool: workerPool,
		Workers:    []*Worker{},
	}
}


func (q *Queue) AddJob(job Job) {
	q.mu.Lock()          
	defer q.mu.Unlock()  

	
	q.Jobs = append(q.Jobs, job)
	fmt.Printf("Job added: %s (ID: %s)\n", job.Name, job.ID)

	
	q.JobChannel <- job
}


func (q *Queue) GetJobs() []Job {
	q.mu.Lock()          
	defer q.mu.Unlock()  
	return q.Jobs        
}


func (q *Queue) RemoveJob(jobID string) error {
	q.mu.Lock()          
	defer q.mu.Unlock()  

	
	for i, job := range q.Jobs {
		if job.ID == jobID {
			
			q.Jobs = append(q.Jobs[:i], q.Jobs[i+1:]...)
			fmt.Printf("Job removed: %s (ID: %s)\n", job.Name, job.ID)
			return nil
		}
	}
	return fmt.Errorf("job with ID %s not found", jobID)
}


func (q *Queue) Start() {
	for i := 1; i <= q.WorkerPool; i++ {
		worker := NewWorker(i, q)
		q.Workers = append(q.Workers, worker)
		worker.Start()
	}
}


func (q *Queue) SubmitJob(name string, timeout int) {
	job := Job{
		ID:      fmt.Sprintf("%d", len(q.Jobs)+1), 
		Name:    name,
		Timeout: timeout, 
	}
	q.AddJob(job)
}
