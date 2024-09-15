package api

import (
	"encoding/json"
	"net/http"
)

// ProblemDetails represents the structure for error details based on RFC 7807
type ProblemDetails struct {
	Type     string `json:"type"`
	Title    string `json:"title"`
	Status   int    `json:"status"`
	Detail   string `json:"detail"`
	Instance string `json:"instance,omitempty"`
}

// WriteProblemResponse writes a problem details response to the client
func WriteProblemResponse(w http.ResponseWriter, problem ProblemDetails) {
	w.Header().Set("Content-Type", "application/problem+json")
	w.WriteHeader(problem.Status)
	json.NewEncoder(w).Encode(problem)
}
