package httpext

import (
	"encoding/json"
	"net/http"

	"github.com/kcchandra/golang-hook/api/response"
)

func Error(w http.ResponseWriter, status int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(response.NewBaseResponse(status, message, nil))
}

func StatusOK(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response.NewBaseResponse(http.StatusOK, "success", data))
}
