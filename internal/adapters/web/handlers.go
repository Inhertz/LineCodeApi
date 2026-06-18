package web

import (
	"LineCodeApi/internal/application"
	"LineCodeApi/internal/core/models"
	"encoding/json"
	"net/http"
)

// getAllManchester handles the /manchester GET endpoint
// and returns all Manchester data from the API or an error if it occurs.
func getAllManchester(api application.APIPort, w http.ResponseWriter, r *http.Request) {
	ans, err := api.GetAllManchester()
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(ans)
	}
}

// generateEncodedManchester handles the /manchester/encoder POST endpoint
// and generates encoded Manchester data from the request body,
// returning the encoded data or an error if it occurs.
func generateEncodedManchester(api application.APIPort, w http.ResponseWriter, r *http.Request) {
	var manchester models.Manchester
	if err := json.NewDecoder(r.Body).Decode(&manchester); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	err := api.GenerateEncodedManchester(&manchester)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(manchester)
	}
}

// generateDecodedManchester handles the /manchester/decoder POST endpoint
// and generates decoded Manchester data from the request body,
// returning the decoded data or an error if it occurs.
func generateDecodedManchester(api application.APIPort, w http.ResponseWriter, r *http.Request) {
	var manchester models.Manchester
	if err := json.NewDecoder(r.Body).Decode(&manchester); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	err := api.GenerateDecodedManchester(&manchester)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(manchester)
	}
}
