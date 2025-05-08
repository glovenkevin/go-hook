package router

import (
	"net/http"

	"github.com/kcchandra/golang-hook/api"
	httpext "github.com/kcchandra/golang-hook/pkg/http"
)

func RegisterRoutes(router *http.ServeMux) {
	router.Handle("/api/health", httpext.RouteWrapper(http.MethodGet, api.HealthCheck))
	router.Handle("/api/ping", httpext.RouteWrapper(http.MethodGet, api.Ping))

	router.Handle("/api/hook/deploy", httpext.RouteWrapper(http.MethodPost, api.DeployHook))
}
