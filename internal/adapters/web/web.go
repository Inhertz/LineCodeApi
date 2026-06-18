package web

import (
	"LineCodeApi/internal/application"
	"context"
	"log"
	"net/http"
	"sync"
	"time"
)

type Adapter struct {
	api  application.APIPort
	port string
}

// NewAdapter creates a new Adapter listening on the given port
func NewAdapter(api application.APIPort, port string) *Adapter {
	return &Adapter{api: api, port: port}
}

// RunAsync runs the server with a wait group for concurrency
func (a Adapter) RunAsync(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	a.Run(ctx)
}

// Run starts the Adapter Web Server and sets up the net/http router.
// It defines the HTTP handlers for the endpoints using the new Go 1.22+ routing.
// Run blocks until ctx is cancelled, then shuts the server down gracefully.
func (a Adapter) Run(ctx context.Context) {
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

	srv := &http.Server{
		Addr:    ":" + a.port,
		Handler: mux,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("failed to serve HTTP server on port %s: %v", a.port, err)
		}
	}()
	log.Printf("HTTP server listening on port %s", a.port)

	<-ctx.Done()

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := srv.Shutdown(shutdownCtx); err != nil {
		log.Printf("HTTP server graceful shutdown failed: %v", err)
		return
	}
	log.Println("HTTP server stopped")
}
