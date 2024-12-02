package handler

import (
	"net/http"

	"github.com/pulsone21/powner/internal/server/response"
)

func setupHandler(fn response.ResponseFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fn(w, r).Respond(w, r)
	})
}
