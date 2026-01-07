package main

import (
	"encoding/json"
	"expvar"
	_ "expvar"
	"flag"
	"fmt"
	"io"
	slog "log/slog"
	"net/http"
	"os"

	"github.com/ardanlabs/nlp"
	"github.com/ardanlabs/nlp/stemmer"
)

var stemCalls = expvar.NewInt("stem.calls")

/* configuration
defaults < configuration file <environment variables <command line
configuration : YAML, TOML... (not stdlib)
environment: os.Getenv
cammand line: flag

external
- viper + cobra = configuration + commandline
- ardanlabs/conf https://github.com/ardanlabs/conf
*/

var config struct {
	Addr string
}

func main() {
	// standard librabry
	config.Addr = os.Getenv("NLP_ADDR")
	if config.Addr == "" {
		config.Addr = ":8080"
	}
	flag.StringVar(&config.Addr, "addr", config.Addr, "address to listen on")
	flag.Parse()

	// IMPORTANT: validate config
	if err := health(); err != nil {
		fmt.Fprintf(os.Stderr, "error: healthcheck -%s\n", err)
		os.Exit(1)
	}
	api := API{
		log: slog.Default().With("app", "nlp"),
	}

	api.log.Info("server starting", "address", config.Addr)
	if err := http.ListenAndServe(config.Addr, nil); err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}

	// Routing
	http.HandleFunc("GET /health", api.healthHandler)
	http.HandleFunc("POST /tokenize", api.tokenizeHandler)
	http.HandleFunc("GET /stem/{word}", api.stemHandler)

	// will look at all interfaces e.g. localhost:8080
	addr := ":8080"
	if err := http.ListenAndServe(addr, nil); err != nil {
		fmt.Fprintf(os.Stderr, "error %s\n", err)
		os.Exit(1)
	}
}

type API struct {
	log *slog.Logger
}

func (a *API) stemHandler(w http.ResponseWriter, r *http.Request) {
	stemCalls.Add(1)
	word := r.PathValue("word")
	a.log.Info("stem", "word", word)
	fmt.Fprintln(w, stemmer.Stem(word))
}

func (a *API) healthHandler(w http.ResponseWriter, r *http.Request) {
	if err := health(); err != nil {
		a.log.Error("health", "error", err)
		http.Error(w, "health check failed", http.StatusInternalServerError)
	}
	fmt.Fprintln(w, "OK")
}

func (a *API) tokenizeHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	b, err := io.ReadAll(r.Body)
	if err != nil {
		a.log.Error("read", "error", err, "remote", r.RemoteAddr)
		http.Error(w, "request body empty", http.StatusBadRequest)
		return
	}
	// validation added by instructor
	if len(string(b)) == 0 {
		a.log.Error("read", "error", "empty request")
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
