package handler

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/pulsone21/powner/internal/server/response"
	"github.com/pulsone21/powner/internal/service"
	"github.com/pulsone21/powner/internal/ui/partials"
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
	t.HandleFunc("GET /teams/list", setupHandler(h.serveTeamList))
	t.HandleFunc("DELETE /teams/{id}", setupHandler(h.deleteTeamRequest))
	t.HandleFunc("GET /teams/{id}/members", setupHandler(h.serveTeamMemberlist))
	t.HandleFunc("GET /teams/{id}/skills", setupHandler(h.serveTeamSkills))
	t.HandleFunc("DELETE /teams/{id}/members/{mID}", setupHandler(h.removeMemberFromTeam))
	t.HandleFunc("GET /teams/{id}/members/{mID}", setupHandler(h.addMemberToTeam))
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
	return response.NewUIResponse(partials.TeamList(*teams), nil)
}

// Path: /partials/teams/{id}/skills
func (h *TeamPartialsHandler) serveTeamSkills(w http.ResponseWriter, r *http.Request) response.IResponse {
	id := r.PathValue("id")
	t, err := h.tServ.GetTeamByID(id)
	if err != nil {
		return response.NewUIResponse(nil, err)
	}

	return response.NewUIResponse(partials.SkillList(t.Skills, t, "Team has no needed skills"), nil)
}
