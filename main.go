package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

type greetingRequest struct {
	Name string `json:"name"`
}

type greetingResponse struct {
	Name      string `json:"name"`
	Message   string `json:"message"`
	Timestamp string `json:"timestamp"`
}

type healthResponse struct {
	Status string `json:"status"`
	App    string `json:"app"`
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	address := ":" + port
	log.Printf("Go Greeting API running at http://localhost%s", address)
	log.Fatal(http.ListenAndServe(address, newRouter()))
}

func newRouter() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", homeHandler)
	mux.HandleFunc("/health", healthHandler)
	mux.HandleFunc("/greet", greetHandler)
	return securityHeaders(mux)
}

func securityHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-Content-Type-Options", "nosniff")
		next.ServeHTTP(w, r)
	})
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	fmt.Fprintln(w, "Welcome to the Go Greeting API. Try /greet?name=Amina or /health.")
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	writeJSON(w, http.StatusOK, healthResponse{
		Status: "ok",
		App:    "go-greeting-api",
	})
}

func greetHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		name := r.URL.Query().Get("name")
		writeJSON(w, http.StatusOK, newGreeting(name))
	case http.MethodPost:
		var request greetingRequest
		decoder := json.NewDecoder(http.MaxBytesReader(w, r.Body, 1_048_576))
		if err := decoder.Decode(&request); err != nil {
			http.Error(w, "invalid JSON body. Example: {\"name\":\"Amina\"}", http.StatusBadRequest)
			return
		}

		writeJSON(w, http.StatusOK, newGreeting(request.Name))
	default:
		w.Header().Set("Allow", "GET, POST")
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

func newGreeting(name string) greetingResponse {
	cleanName := strings.Join(strings.Fields(name), " ")
	if cleanName == "" {
		cleanName = "World"
	}

	return greetingResponse{
		Name:      cleanName,
		Message:   fmt.Sprintf("Hello, %s! Welcome to Go.", cleanName),
		Timestamp: time.Now().UTC().Format(time.RFC3339),
	}
}

func writeJSON(w http.ResponseWriter, status int, payload any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(payload); err != nil {
		log.Printf("could not write JSON response: %v", err)
	}
}
