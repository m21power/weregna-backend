package utils

import (
	"encoding/json"
	"net/http"
	"time"
)

// StandardResponse is the base structure for all API responses
type StandardResponse struct {
	Success   bool        `json:"success"`
	Data      interface{} `json:"data,omitempty"`
	Error     string      `json:"error,omitempty"`
	Message   string      `json:"message,omitempty"`
	Timestamp string      `json:"timestamp"`
}

// WriteJSON writes a standardized JSON response with proper headers
func WriteJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	response := StandardResponse{
		Success:   status >= http.StatusOK && status < http.StatusBadRequest,
		Data:      data,
		Timestamp: time.Now().UTC().Format(time.RFC3339),
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		// Fallback to simple error response if JSON encoding fails
		http.Error(w, `{"success":false,"error":"failed to encode response"}`, http.StatusInternalServerError)
	}
}

// WriteError writes a standardized error response
func WriteError(w http.ResponseWriter, err error, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	response := StandardResponse{
		Success:   false,
		Error:     err.Error(),
		Timestamp: time.Now().UTC().Format(time.RFC3339),
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, `{"success":false,"error":"failed to encode error"}`, http.StatusInternalServerError)
	}
}

// WriteSuccess writes a success message with optional data
func WriteSuccess(w http.ResponseWriter, message string, data interface{}, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	response := StandardResponse{
		Success:   true,
		Message:   message,
		Data:      data,
		Timestamp: time.Now().UTC().Format(time.RFC3339),
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, `{"success":false,"error":"failed to encode response"}`, http.StatusInternalServerError)
	}
}
