package main

import (
	"fmt"
	"goqueue/internal"
	"log"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"
	"os"
)

func main() {
	queue := internal.NewQueue(3)
	queue.Start()


	backendRouter := mux.NewRouter()

	
	cors := handlers.CORS(
		handlers.AllowedOrigins([]string{"http://localhost:8081"}), 
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"}),
		handlers.AllowedHeaders([]string{"Content-Type"}),
	)

	backendRouter.HandleFunc("/submit", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Received request on /submit")
		var jobData struct {
			Name    string `json:"name"`
			Timeout int    `json:"timeout"`
		}

		if err := json.NewDecoder(r.Body).Decode(&jobData); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}


		queue.SubmitJob(jobData.Name, jobData.Timeout)

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{
			"message": fmt.Sprintf("Job '%s' submitted successfully!", jobData.Name),
		})
	}).Methods("POST")


	go func() {
		log.Println("Starting backend server on http://localhost:8080")
		if err := http.ListenAndServe(":8080", cors(backendRouter)); err != nil {
			log.Fatal("Backend server failed to start:", err)
		}
	}()


	frontendRouter := mux.NewRouter()


	frontendDir := "./web" 
	if _, err := os.Stat(frontendDir); os.IsNotExist(err) {
		log.Fatal("Frontend directory not found, please ensure the 'web' folder exists")
	}
	frontendRouter.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir(frontendDir))))

	log.Println("Starting frontend server on http://localhost:8081")
	if err := http.ListenAndServe(":8081", frontendRouter); err != nil {
		log.Fatal("Frontend server failed to start:", err)
	}
}
