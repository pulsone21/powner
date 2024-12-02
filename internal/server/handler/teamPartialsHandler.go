package handler

import (
	"net/http"

	"github.com/pulsone21/powner/internal/server/middleware"
	"github.com/pulsone21/powner/internal/server/response"
	"github.com/pulsone21/powner/internal/service"
	"github.com/pulsone21/powner/internal/ui/subpage"
)

type TeamPartialsHandler struct {
	tServ service.TeamService
}

func NewTeamPartialsHandler(tServ service.TeamService) *TeamPartialsHandler {
	return &TeamPartialsHandler{
		tServ: tServ,
	}
}

func (h *TeamPartialsHandler) RegisterRoutes(t *http.ServeMux) {
	t.HandleFunc("/teams/overview", setupHandler(h.serveTeamsOverview))
}

func (h *TeamPartialsHandler) serveTeamsOverview(w http.ResponseWriter, r *http.Request) response.IResponse {
	log := middleware.GetLogger(r.Context())
	log.Debug("team overview partial requested")
	t, err := h.tServ.GetTeams()
	if err != nil {
		return nil
	}
	log.Debug("found all teams")
	return response.NewUIResponse(subpage.TeamsOverview(*t), nil)
}
