package handler

import (
	"fmt"
	"net/http"

	"github.com/a-h/templ"
	"github.com/pulsone21/powner/internal/server/middleware"
	"github.com/pulsone21/powner/internal/server/response"
	"github.com/pulsone21/powner/internal/service"
	"github.com/pulsone21/powner/internal/ui/components"
	"github.com/pulsone21/powner/internal/ui/partials"
	"github.com/pulsone21/powner/internal/ui/subpage"
)

type TeamPartialsHandler struct {
	tServ service.TeamService
	mServ service.MemberManagementService
}

func NewTeamPartialsHandler(tServ service.TeamService, mServ service.MemberManagementService) *TeamPartialsHandler {
	return &TeamPartialsHandler{
		tServ: tServ,
		mServ: mServ,
	}
}

func (h *TeamPartialsHandler) RegisterRoutes(t *http.ServeMux) {
	t.HandleFunc("GET /teams/overview", setupHandler(h.serveTeamsOverview))
	t.HandleFunc("GET /teams/list", setupHandler(h.serveTeamList))
	t.HandleFunc("DELETE /teams/{id}", setupHandler(h.deleteTeamRequest))
	t.HandleFunc("GET /teams/{id}/details", setupHandler(h.serveTeamsDetails))
	t.HandleFunc("GET /teams/{id}/members", setupHandler(h.serveTeamMemberlist))
	t.HandleFunc("DELETE /teams/{id}/members/{mID}", setupHandler(h.removeMemberFromTeam))
	t.HandleFunc("GET /teams/{id}/members/{mID}", setupHandler(h.addMemberToTeam))
}

func (h *TeamPartialsHandler) serveTeamsOverview(w http.ResponseWriter, r *http.Request) response.IResponse {
	log := middleware.GetLogger(r.Context())
	log.Debug("team overview partial requested")

	t, err := h.tServ.GetTeams()
	if err != nil {
		return response.NewUIResponse(nil, err)
	}

	log.Debug("found all teams")
	return response.NewUIResponse(subpage.TeamsOverview(*t, nil), nil)
}

func (h *TeamPartialsHandler) serveTeamsDetails(w http.ResponseWriter, r *http.Request) response.IResponse {
	id := r.PathValue("id")
	t, err := h.tServ.GetTeamByID(id)
	if err != nil {
		return response.NewUIResponse(nil, err)
	}

	if t == nil {
		return response.NewUIResponse(nil, fmt.Errorf("Couldn't find team with id: %v", id))
	}

	return response.NewUIResponse(subpage.TeamDetails(*t), nil)
}

func (h *TeamPartialsHandler) deleteTeamRequest(w http.ResponseWriter, r *http.Request) response.IResponse {
	id := r.PathValue("id")
	err := h.tServ.DeleteTeam(id)
	if err != nil {
		return response.NewUIResponse(nil, err)
	}

	w.Header().Add("HX-Trigger", service.DeleteTeamEvent)
	return response.NewUIResponse(templ.NopComponent, nil)
}

func (h *TeamPartialsHandler) serveTeamMemberlist(w http.ResponseWriter, r *http.Request) response.IResponse {
	id := r.PathValue("id")
	t, err := h.tServ.GetTeamByID(id)
	if err != nil {
		return response.NewUIResponse(nil, err)
	}

	return response.NewUIResponse(partials.TeamMemberList(*t), nil)
}

func (h *TeamPartialsHandler) removeMemberFromTeam(w http.ResponseWriter, r *http.Request) response.IResponse {
	id := r.PathValue("id")
	mID := r.PathValue("mID")

	t, err := h.mServ.RemoveMemberToTeam(id, mID)
	if err != nil {
		return response.NewUIResponse(nil, err)
	}

	w.Header().Add("HX-Trigger", service.ChangeTeamEvent)
	return response.NewUIResponse(partials.TeamMemberList(*t), nil)
}

func (h *TeamPartialsHandler) addMemberToTeam(w http.ResponseWriter, r *http.Request) response.IResponse {
	id := r.PathValue("id")
	mID := r.PathValue("mID")

	t, err := h.mServ.AddMemberToTeam(id, mID)
	if err != nil {
		return response.NewUIResponse(nil, err)
	}

	w.Header().Add("HX-Trigger", service.ChangeTeamEvent)
	return response.NewUIResponse(partials.TeamMemberList(*t), nil)
}

// Path: /partials/teams/list
func (h *TeamPartialsHandler) serveTeamList(w http.ResponseWriter, r *http.Request) response.IResponse {
	teams, err := h.tServ.GetTeams()
	if err != nil {
		return response.NewUIResponse(nil, err)
	}
	return response.NewUIResponse(partials.TeamList(*teams, "No teams found", components.DeleteTeamButton), nil)
}
