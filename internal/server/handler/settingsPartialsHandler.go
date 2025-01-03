package handler

import (
	"net/http"

	"github.com/pulsone21/powner/internal/server/response"
	"github.com/pulsone21/powner/internal/ui/subpage"
)

type SettingsPartialsHandler struct{}

func NewSettingsPartialsHandler() *SettingsPartialsHandler {
	return &SettingsPartialsHandler{}
}

func (h *SettingsPartialsHandler) RegisterRoutes(t *http.ServeMux) {
	t.HandleFunc("GET /settings/overview", setupHandler(h.serveSettingsOverview))
}

func (h *SettingsPartialsHandler) serveSettingsOverview(w http.ResponseWriter, r *http.Request) response.IResponse {
	return response.NewUIResponse(subpage.SettingsSubpage(), nil)
}
