package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// Command Injection Endpoint
	r.HandleFunc("/execute", executeHandler)

	// SSRF Endpoint
	r.HandleFunc("/fetch", fetchHandler)

	// Environment Variable Disclosure Endpoint
	r.HandleFunc("/env", envHandler)

	// Health check
	r.HandleFunc("/", homeHandler)

	fmt.Println("Server starting on :8080")
	http.ListenAndServe(":8080", r)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Vulnerable Supply Chain Demo Application\n")
	fmt.Fprintf(w, "Endpoints:\n")
	fmt.Fprintf(w, "  /execute?cmd=<command> - Execute OS command (VULNERABLE)\n")
	fmt.Fprintf(w, "  /fetch?url=<url> - Fetch URL (VULNERABLE)\n")
	fmt.Fprintf(w, "  /env - Show environment variables\n")
}

func executeHandler(w http.ResponseWriter, r *http.Request) {
	cmd := r.URL.Query().Get("cmd")
	if cmd == "" {
		http.Error(w, "Command parameter is required", http.StatusBadRequest)
		return
	}

	// VULNERABLE: Direct command execution without sanitization
	out, err := exec.Command("/bin/sh", "-c", cmd).CombinedOutput()
	if err != nil {
		http.Error(w, fmt.Sprintf("Error: %s, Output: %s", err, string(out)), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintf(w, "Command: %s\nOutput:\n%s", cmd, string(out))
}

func fetchHandler(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Query().Get("url")
	if url == "" {
		http.Error(w, "URL parameter is required", http.StatusBadRequest)
		return
	}

	// VULNERABLE: No validation of the URL before fetching
	resp, err := http.Get(url)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error fetching URL: %s", err), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error reading response: %s", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintf(w, "URL: %s\nStatus: %s\nResponse:\n%s", url, resp.Status, string(body))
}

func envHandler(w http.ResponseWriter, r *http.Request) {
	envVars := os.Environ()

	w.Header().Set("Content-Type", "text/plain")
	for _, envVar := range envVars {
		fmt.Fprintln(w, envVar)
	}
}
