package router

import "net/http"

func NewPartialsRouter(handler ...IRouter) *http.ServeMux {
	// f := http.NewServeMux()
	t := http.NewServeMux()
	for _, h := range handler {
		h := h
		h.RegisterRoutes(t)
	}

	t.HandleFunc("GET /", serveNotFound())
	return t
}
