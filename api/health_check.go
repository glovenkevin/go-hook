package api

import (
	"net/http"

	httpext "github.com/kcchandra/golang-hook/pkg/http"
)

// HealthCheck handles the health check endpoint
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	httpext.StatusOK(w, "OK")
}
