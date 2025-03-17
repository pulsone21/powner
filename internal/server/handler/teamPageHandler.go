package handler

import (
	"net/http"

	"github.com/pulsone21/powner/internal/server/middleware"
	"github.com/pulsone21/powner/internal/server/response"
	"github.com/pulsone21/powner/internal/service"
	"github.com/pulsone21/powner/internal/ui/pages"
	"github.com/pulsone21/powner/internal/ui/subpage"
)

type TeamPageHandler struct {
	teamService service.TeamService
}

func NewTeamPageHandler(ser service.TeamService) TeamPageHandler {
	return TeamPageHandler{
		teamService: ser,
	}
}

func (h TeamPageHandler) GetRoutes() *http.ServeMux {
	t := http.NewServeMux()
	t.HandleFunc("GET /", setupHandler(h.generalTeamPage))
	t.HandleFunc("GET /{id}", setupHandler(h.specificTeamPage))
	return t
}

func (h *TeamPageHandler) generalTeamPage(w http.ResponseWriter, r *http.Request) response.IResponse {
	log := middleware.GetLogger(r.Context())
	log.Debug("general team page requested")
	t, err := h.teamService.GetTeams()
	if err != nil {
		return response.NewUIResponse(nil, err)
	}

	if ok := r.Header.Get("Hx-Request"); ok != "" {
		return response.NewUIResponse(subpage.TeamsOverview(*t, true), nil)
	}

	return response.NewUIResponse(pages.TeamsOverviewPage(*t), nil)
}

func (h *TeamPageHandler) specificTeamPage(w http.ResponseWriter, r *http.Request) response.IResponse {
	log := middleware.GetLogger(r.Context())
	log.Debug("specific team page requested")
	t, err := h.teamService.GetTeamByID(r.PathValue("id"))
	if err != nil {
		return response.NewUIResponse(nil, err)
	}

	if ok := r.Header.Get("Hx-Request"); ok != "" {
		return response.NewUIResponse(subpage.TeamDetails(*t, true), nil)
	}

	return response.NewUIResponse(pages.TeamDetailPage(*t), nil)
}
