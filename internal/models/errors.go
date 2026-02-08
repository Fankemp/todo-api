package models

import (
	"encoding/json"
	"net/http"
)

type ResponseError struct {
	Message string `json:"error"`
}

func SendError(w http.ResponseWriter, message string, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	errResp := ResponseError{Message: message}
	err := json.NewEncoder(w).Encode(errResp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
