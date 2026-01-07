package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/ardanlabs/nlp"
)

func main() {
	// Routing
	http.HandleFunc("GET /health", healthHandler)
	http.HandleFunc("POST /tokenize", tokenizeHandler)

	// will look at all interfaces e.g. localhost:8080
	addr := ":8080"
	if err := http.ListenAndServe(addr, nil); err != nil {
		fmt.Fprintf(os.Stderr, "error %s\n", err)
		os.Exit(1)
	}
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	if err := health(); err != nil {
		http.Error(w, "health check failed", http.StatusInternalServerError)
	}
	fmt.Fprintln(w, "OK")
}

func tokenizeHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	b, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "request body empty", http.StatusNoContent)
	}
	tokens := nlp.Tokenize(string(b))
	t := map[string]any{"tokens": tokens}
	js, err := json.Marshal(t)
	if err != nil {
		http.Error(w, "request body empty", http.StatusInternalServerError)
	}
	fmt.Fprintln(w, string(js))
}

func health() error {
	// TODO: Actual health check
	return nil
}
