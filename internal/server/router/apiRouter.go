package router

import (
	"fmt"
	"net/http"
)

// Creates for all given handlers the routes to the path /api with the given version number. Panics if you give it a lower number then 1
func NewApiRouter(version int, handler ...IRouter) *http.ServeMux {
	if version < 1 {
		panic("Api version must be higher then 0")
	}

	api := http.NewServeMux()
	for _, h := range handler {
		api.Handle(
			fmt.Sprintf("/api/v%b/", version),
			http.StripPrefix(fmt.Sprintf("/api/v%b", version), h.GetRoutes()))
	}
	return api
}
