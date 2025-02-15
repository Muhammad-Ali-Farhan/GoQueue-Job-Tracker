package api

import (
	"fmt"
	"goqueue/internal"
	"log"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
)

func StartServer(queue *internal.Queue) {
	router := mux.NewRouter()

	
	router.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("web"))))

	router.HandleFunc("/submit", func(w http.ResponseWriter, r *http.Request) {
		var jobData struct {
			Name     string `json:"name"`
			Priority int    `json:"priority"`
		}

		if err := json.NewDecoder(r.Body).Decode(&jobData); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		queue.SubmitJob(jobData.Name, jobData.Priority)

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{
			"message": fmt.Sprintf("Job '%s' submitted successfully with priority %d!", jobData.Name, jobData.Priority),
		})
	}).Methods("POST")

	
	router.HandleFunc("/status/{jobID}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		jobID := vars["jobID"]

		status := queue.JobStatus(jobID)
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{
			"status": status,
		})
	}).Methods("GET")

	log.Println("Starting API server on http://localhost:8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
