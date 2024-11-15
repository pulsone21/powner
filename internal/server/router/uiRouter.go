package router

import "net/http"

func NewFrontendRouter(uiHandler IRouter) *http.ServeMux {
	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("./public/static/"))

	mux.Handle("/static/", http.StripPrefix("/static/", fs))
	mux.Handle("/", uiHandler.GetRoutes())

	return mux
}
