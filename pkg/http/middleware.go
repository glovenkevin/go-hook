package httpext

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/kcchandra/golang-hook/api/response"
	"github.com/kcchandra/golang-hook/pkg/constant"
)

func MiddlewareWrapper(next http.HandlerFunc) http.HandlerFunc {
	return LoggingMiddleware(RecoveryMiddleware(next))
}

// LoggingMiddleware logs request details
func LoggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now().UTC()

		next.ServeHTTP(w, r)

		log.Printf(
			"%s %s %s %v",
			r.Method,
			r.RequestURI,
			r.RemoteAddr,
			time.Since(start),
		)
	}
}

// RecoveryMiddleware recovers from panics
func RecoveryMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("panic: %v", err)
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(response.NewBaseResponse(
					http.StatusInternalServerError,
					constant.ErrorInternalServerError,
					nil,
				))
			}
		}()

		next.ServeHTTP(w, r)
	}
}
