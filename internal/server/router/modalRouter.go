package router

import "net/http"

func NewModalRouter(handler ...IRouter) *http.ServeMux {
	t := http.NewServeMux()
	for _, h := range handler {
		h := h
		h.RegisterRoutes(t)
	}

	t.HandleFunc("GET /", serveNotFound())
	return t
}
