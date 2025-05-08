package httpext

import (
	"net/http"
)

type Route struct {
	Path    string
	Handler http.HandlerFunc
}

func RouteWrapper(method string, handler http.HandlerFunc) http.HandlerFunc {
	return MiddlewareWrapper(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != method {
			Error(w, http.StatusMethodNotAllowed, "Method not allowed")
			return
		}

		handler.ServeHTTP(w, r)
	})
}
