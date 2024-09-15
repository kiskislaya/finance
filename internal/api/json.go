package api

import (
	"encoding/json"
	"net/http"
)

// WriteJSONResponse writes a JSON response to the client with the given status code and data
func WriteJSONResponse(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		// In case of JSON encoding error, write a generic internal server error
		WriteProblemResponse(w, ProblemDetails{
			Status: http.StatusInternalServerError,
		})
	}
}
