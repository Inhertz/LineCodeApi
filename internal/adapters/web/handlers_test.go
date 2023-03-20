package web

import (
	"LineCodeApi/internal/core/models"
	"LineCodeApi/test/mocks"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetAllManchester(t *testing.T) {

	expectedResponse := `[{"id":1,"decoded":"0AE3","encoded":"-A+A-A+A-A+A-A+A+A-A-A+A+A-A-A+A+A-A+A-A+A-A-A+A-A+A-A+A+A-A+A-A","decodedPulseWidth":400,"encodedPulseWidth":200,"unit":"us"},{"id":2,"decoded":"7e","encoded":"-A+A+A-A+A-A+A-A+A-A+A-A+A-A-A+A","decodedPulseWidth":800,"encodedPulseWidth":400,"unit":"us"}]`
	var expOut []models.Manchester
	json.Unmarshal([]byte(expectedResponse), &expOut)
	mockAPI := &mocks.ApplicationMock{ExpectedOutput: expOut}

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	getAllManchester(mockAPI, c)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, expectedResponse, w.Body.String())

}

func TestGenerateEncodedManchester(t *testing.T) {
	requestBody := `{"decoded":"0AE3","decodedPulseWidth":400}`
	expectedResponse := `{"id":1,"decoded":"0AE3","encoded":"-A+A-A+A-A+A-A+A+A-A-A+A+A-A-A+A+A-A+A-A+A-A-A+A-A+A-A+A+A-A+A-A","decodedPulseWidth":400,"encodedPulseWidth":200,"unit":"us"}`

	var expOut models.Manchester
	err := json.Unmarshal([]byte(expectedResponse), &expOut)
	if err != nil {
		t.Fatal(err)
	}
	mockAPI := &mocks.ApplicationMock{ExpectedOutput: expOut}

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest(http.MethodPost, "/manchester/encoder", bytes.NewBufferString(requestBody))
	c.Request.Header.Set("Content-Type", "application/json")

	generateEncodedManchester(mockAPI, c)

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Equal(t, expectedResponse, w.Body.String())
}

func TestGenerateDecodedManchester(t *testing.T) {
	requestBody := `{"encoded":"-A+A+A-A+A-A+A-A+A-A+A-A+A-A-A+A","encodedPulseWidth":400,"unit":"us"}`
	expectedResponse := `{"id":2,"decoded":"7e","encoded":"-A+A+A-A+A-A+A-A+A-A+A-A+A-A-A+A","decodedPulseWidth":800,"encodedPulseWidth":400,"unit":"us"}`

	var expOut models.Manchester
	err := json.Unmarshal([]byte(expectedResponse), &expOut)
	if err != nil {
		t.Fatal(err)
	}
	mockAPI := &mocks.ApplicationMock{ExpectedOutput: expOut}

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest(http.MethodPost, "/manchester/decoder", bytes.NewBufferString(requestBody))
	c.Request.Header.Set("Content-Type", "application/json")

	generateDecodedManchester(mockAPI, c)

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Equal(t, expectedResponse, w.Body.String())
}
