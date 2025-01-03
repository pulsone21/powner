package handler

import (
	"net/http"

	"github.com/pulsone21/powner/internal/server/response"
	"github.com/pulsone21/powner/internal/ui/pages"
)

type SettingsPageHandler struct{}

func NewSettingsPageHandler() *SettingsPageHandler {
	return &SettingsPageHandler{}
}

func (h SettingsPageHandler) GetRoutes() *http.ServeMux {
	t := http.NewServeMux()
	t.HandleFunc("GET /", setupHandler(h.serveSettingsPage))
	return t
}

func (h *SettingsPageHandler) serveSettingsPage(w http.ResponseWriter, r *http.Request) response.IResponse {
	return response.NewUIResponse(pages.Settings(), nil)
}
