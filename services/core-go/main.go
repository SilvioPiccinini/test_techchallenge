package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type HealthResponse struct {
	Status string `json:"status"`
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	// Apenas GET
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	// Headers de segurança
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.Header().Set("Cache-Control", "no-cache")

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(HealthResponse{Status: "ok"}); err != nil {
		log.Printf("erro ao encodar resposta de health check: %v", err)
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/healthz", healthCheck)

	server := &http.Server{
		Addr:           ":8080",
		Handler:        mux,
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   5 * time.Second,
		IdleTimeout:    30 * time.Second,
	}

	log.Println("Servidor rodando em :8080")
	log.Fatal(server.ListenAndServe())
}