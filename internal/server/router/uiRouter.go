package router

import "net/http"

func NewFrontendRouter(api http.Handler, uiHandler IRouter) *http.ServeMux {
	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("./public/static/"))

	mux.Handle("/static/", http.StripPrefix("/static/", fs))
	mux.Handle("/api/", http.StripPrefix("/api", api))

	mux.Handle("/", uiHandler.GetRoutes())

	return mux
}
