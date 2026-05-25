package util

import (
	"encoding/json"
	"net/http"
)

func SendError(w http.ResponseWriter, statusCode int, msg string) {
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(msg)
}
