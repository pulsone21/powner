package handler

import (
	"net/http"

	"github.com/pulsone21/powner/internal/server/response"
	"github.com/pulsone21/powner/internal/ui/pages"
)

type GeneralPageHandler struct{}

func NewGeneralPageHandler() *GeneralPageHandler {
	return &GeneralPageHandler{}
}

func (h GeneralPageHandler) GetRoutes() *http.ServeMux {
	t := http.NewServeMux()
	t.HandleFunc("GET /settings", setupHandler(h.serveSettingsPage))
	t.HandleFunc("GET /{$}", setupHandler(h.serveIndexPage))
	t.HandleFunc("GET /", setupHandler(h.serveNotFound))
	return t
}

func (h *GeneralPageHandler) serveIndexPage(w http.ResponseWriter, r *http.Request) response.IResponse {
	return response.NewUIResponse(
		pages.Index(), nil)
}

func (h *GeneralPageHandler) serveNotFound(w http.ResponseWriter, r *http.Request) response.IResponse {
	return response.NewUIResponse(pages.NotFound(), nil)
}

func (h *GeneralPageHandler) serveSettingsPage(w http.ResponseWriter, r *http.Request) response.IResponse {
	return response.NewUIResponse(pages.Settings(), nil)
}
