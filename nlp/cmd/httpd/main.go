package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/ardanlabs/nlp"
	"github.com/ardanlabs/nlp/stemmer"
)

func main() {
	// Routing
	http.HandleFunc("GET /health", healthHandler)
	http.HandleFunc("POST /tokenize", tokenizeHandler)
	http.HandleFunc("GET /stem/{word}", stemHandler)

	// will look at all interfaces e.g. localhost:8080
	addr := ":8080"
	if err := http.ListenAndServe(addr, nil); err != nil {
		fmt.Fprintf(os.Stderr, "error %s\n", err)
		os.Exit(1)
	}
}

func stemHandler(w http.ResponseWriter, r *http.Request) {
	word := r.PathValue("word")
	fmt.Fprintln(w, stemmer.Stem(word))
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
		http.Error(w, "request body empty", http.StatusBadRequest)
		return
	}
	// validation added by instructor
	if len(string(b)) == 0 {
		http.Error(w, "empty request", http.StatusBadRequest)
		return
	}

	tokens := nlp.Tokenize(string(b))
	t := map[string]any{"tokens": tokens}
	js, err := json.Marshal(t)
	if err != nil {
		http.Error(w, "request body empty", http.StatusInternalServerError)
	}
	fmt.Fprintln(w, string(js))

	// Instructors code
	// w.Header().Set("content-type", "application/json")
	// resp := map[string]any{"tokens": tokens}
	// json.NewEncoder(w).Encode(resp)
}

// Step 1 Readdata, parse, *validate*
func health() error {
	// TODO: Actual health check
	return nil
}
