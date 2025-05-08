package api

import (
	"net/http"

	httpext "github.com/kcchandra/golang-hook/pkg/http"
)

func Ping(w http.ResponseWriter, r *http.Request) {
	httpext.StatusOK(w, "pong")
}
