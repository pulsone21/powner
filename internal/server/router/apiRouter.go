package router

import (
	"fmt"
	"net/http"
)

// Creates for all given handlers the routes to the path /api with the given version number.
// Panics if you give it a lower number then 1
func NewApiRouter(version int, handler ...IRouter) *http.ServeMux {
	if version < 1 {
		panic("Api version must be higher then 0")
	}

	apiRoute := fmt.Sprintf("/v%b/", version)

	api := http.NewServeMux()
	t := http.NewServeMux()
	for _, h := range handler {
		h := h
		h.RegisterRoutes(t)
	}

	api.Handle(
		apiRoute,
		http.StripPrefix(apiRoute[:len(apiRoute)-1], t))
	return api
}
