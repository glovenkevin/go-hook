package api

import (
	"encoding/json"
	"net/http"

	"github.com/kcchandra/golang-hook/api/response"
)

func DeployHook(w http.ResponseWriter, r *http.Request) {
	response := response.NewBaseResponse(http.StatusOK, "OK", "success")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
