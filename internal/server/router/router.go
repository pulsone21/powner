package router

import (
	"net/http"

	"github.com/pulsone21/powner/internal/server/response"
	"github.com/pulsone21/powner/internal/ui/pages"
)

type IRouter interface {
	RegisterRoutes(*http.ServeMux)
}

func serveNotFound() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response.NewUIResponse(pages.NotFound(), nil).Respond(w, r)
	})
}
