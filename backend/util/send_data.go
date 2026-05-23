package util

import (
	"encoding/json"
	"net/http"
)

func SendData(w http.ResponseWriter, data interface{}, code int) {
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(data)
}
