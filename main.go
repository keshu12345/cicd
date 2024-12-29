package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

var (
	Version     = "v1.0.0"
	HealthyFlag = true
)

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	if HealthyFlag {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Unhealthy"))
	}
}

func VersionHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Version: %s", Version)))
}

// ToggleHealth simulates a failure or recovery for testing rollback
func ToggleHealth(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	state := query.Get("state")
	if state == "healthy" {
		HealthyFlag = true
		log.Println("Application state set to healthy")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Application state: healthy"))
	} else if state == "unhealthy" {
		HealthyFlag = false
		log.Println("Application state set to unhealthy")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Application state: unhealthy"))
	} else {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid state. Use ?state=healthy or ?state=unhealthy"))
	}
}

func main() {

	http.HandleFunc("/health", HealthHandler)
	http.HandleFunc("/version", VersionHandler)
	http.HandleFunc("/toggle-health", ToggleHealth)

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Starting application version %s on port %s\n", Version, port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
