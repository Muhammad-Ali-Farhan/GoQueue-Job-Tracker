# GoQueue – A Minimalist Distributed Task Queue
# Conceptual Model Project

A lightweight distributed task queue in Go. It uses Goroutines and channels to process jobs asynchronously with fault tolerance, optional persistence, and a simple REST API for submitting jobs.

## Features
- Job queue with worker pool
- Fault tolerance (jobs retry if they fail)
- Persistence (saves jobs to a file)
- Simple CLI and REST API for job submission

## How to Run
1. Clone the repository.
2. Initialize Go module with `go mod init goqueue`.
3. Start the queue with `go run cmd/main.go`.

### API Endpoints
- **POST /submit**: Submit a job to the queue.
    - Request body: `{ "name": "Job name" }`
    - Response: `{ "message": "Job submitted", "job_id": "123" }`

## Dependencies
- Go 1.20
- SQLite (for optional persistence)
