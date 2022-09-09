package web

import (
	"LineCodeApi/internal/application"
	"LineCodeApi/internal/core/models"
	"net/http"
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

// Run will start the Adapter Web Server and listen
func (a Adapter) Run() {
	r := gin.Default()

	r.GET("/manchester", func(c *gin.Context) {
		ans, err := a.api.GetAllManchester()
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusOK, ans)
		}

	})

	r.POST("/manchester/encoder", func(c *gin.Context) {
		var manchester models.Manchester
		c.BindJSON(&manchester)
		err := a.api.GenerateEncodedManchester(&manchester)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusCreated, manchester)
		}
	})

	r.POST("/manchester/decoder", func(c *gin.Context) {
		var manchester models.Manchester
		c.BindJSON(&manchester)
		err := a.api.GenerateDecodedManchester(&manchester)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusCreated, manchester)
		}
	})

	r.Run()
}
