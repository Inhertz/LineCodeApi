package web

import (
	"LineCodeApi/internal/application"
	"LineCodeApi/internal/core/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// getAllManchester handles the /manchester GET endpoint
// and returns all Manchester data from the API or an error if it occurs.
func getAllManchester(api application.APIPort, c *gin.Context) {
	ans, err := api.GetAllManchester()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, ans)
	}
}

// generateEncodedManchester handles the /manchester/encoder POST endpoint
// and generates encoded Manchester data from the request body,
// returning the encoded data or an error if it occurs.
func generateEncodedManchester(api application.APIPort, c *gin.Context) {
	var manchester models.Manchester
	c.BindJSON(&manchester)
	err := api.GenerateEncodedManchester(&manchester)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusCreated, manchester)
	}
}

// generateDecodedManchester handles the /manchester/decoder POST endpoint
// and generates decoded Manchester data from the request body,
// returning the decoded data or an error if it occurs.
func generateDecodedManchester(api application.APIPort, c *gin.Context) {
	var manchester models.Manchester
	c.BindJSON(&manchester)
	err := api.GenerateDecodedManchester(&manchester)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusCreated, manchester)
	}
}
