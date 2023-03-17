package web

import (
	"LineCodeApi/internal/application"
	"sync"

	"github.com/gin-gonic/gin"
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

// Run starts the Adapter Web Server and sets up the Gin router.
// It defines the HTTP handlers for the endpoints, and delegates the application logic
// to separate functions which take the api instance and gin.Context as input.
func (a Adapter) Run() {
	r := gin.Default()

	r.GET("/manchester", func(c *gin.Context) {
		getAllManchester(a.api, c)
	})

	r.POST("/manchester/encoder", func(c *gin.Context) {
		generateEncodedManchester(a.api, c)
	})

	r.POST("/manchester/decoder", func(c *gin.Context) {
		generateDecodedManchester(a.api, c)
	})

	r.Run()
}
