package api

import (
	"encoding/json"
	"net/http"
)

// ErrorResponse standardizes the error message structure
type ErrorResponse struct {
	Error string `json:"error"`
}

// SuccessResponse standardizes the success response structure
type SuccessResponse struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data,omitempty"`
}

// respondWithError sends an error response in JSON format
func respondWithError(w http.ResponseWriter, code int, message string) {
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(ErrorResponse{Error: message})
}

// respondWithSuccess sends a success response in JSON format
func respondWithSuccess(w http.ResponseWriter, data interface{}) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(SuccessResponse{Status: "success", Data: data})
}
