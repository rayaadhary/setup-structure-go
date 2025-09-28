package main

import (
	"encoding/json"
	"net/http"
)

type APIResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func writeJSON(w http.ResponseWriter, status int, resp APIResponse) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(resp)
}

func writeError(w http.ResponseWriter, status int, msg string) {
	writeJSON(w, status, APIResponse{
		Success: false,
		Message: msg,
	})
}

func writeSuccess(w http.ResponseWriter, status int, data interface{}, msg string) {
	writeJSON(w, status, APIResponse{
		Success: true,
		Message: msg,
		Data:    data,
	})
}
