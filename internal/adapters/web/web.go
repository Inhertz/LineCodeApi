package web

import (
	"LineCodeApi/internal/application"
	"net/http"
	"sync"
)

type Adapter struct {
	api application.APIPort
}

// NewAdapter creates a new Adapter
func NewAdapter(api application.APIPort) *Adapter {
	return &Adapter{api: api}
}

// RunAsync runs the server with a wait group for concurrency
func (a Adapter) RunAsync(wg *sync.WaitGroup) {

	a.Run()

	wg.Done()
}

// Run starts the Adapter Web Server and sets up the net/http router.
// It defines the HTTP handlers for the endpoints using the new Go 1.22+ routing.
func (a Adapter) Run() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /manchester", func(w http.ResponseWriter, r *http.Request) {
		getAllManchester(a.api, w, r)
	})

	mux.HandleFunc("POST /manchester/encoder", func(w http.ResponseWriter, r *http.Request) {
		generateEncodedManchester(a.api, w, r)
	})

	mux.HandleFunc("POST /manchester/decoder", func(w http.ResponseWriter, r *http.Request) {
		generateDecodedManchester(a.api, w, r)
	})

	http.ListenAndServe(":8080", mux)
}
