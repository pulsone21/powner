package handler

import (
	"net/http"

	"github.com/pulsone21/powner/internal/server/middleware"
	"github.com/pulsone21/powner/internal/server/response"
	"github.com/pulsone21/powner/internal/service"
)

type MemberManagementHandler struct {
	service service.MemberManagementService
}

func (h MemberManagementHandler) RegisterRoutes(t *http.ServeMux) {
	t.HandleFunc("POST /member/{mem_id}/team/{team_id}", setupApiHandler(h.addMemberToTeam))
	t.HandleFunc("DELETE /member/{mem_id}/team/{team_id}", setupApiHandler(h.removeMemberFromTeam))
}

func (h MemberManagementHandler) addMemberToTeam(w http.ResponseWriter, r *http.Request) response.IResponse {
	log := middleware.GetLogger(r.Context())
	log.Debug("Add member to team hit")

	t, err := h.service.AddMemberToTeam(r.PathValue("team_id"), r.PathValue("mem_id"))
	if err != nil {
		return response.NewApiResponse(nil, err)
	}

	log.Debug("Member added to team")
	return response.NewApiResponse(t, nil)
}

func (h MemberManagementHandler) removeMemberFromTeam(w http.ResponseWriter, r *http.Request) response.IResponse {
	log := middleware.GetLogger(r.Context())
	log.Debug("Remove member from team hit")

	t, err := h.service.RemoveMemberToTeam(r.PathValue("team_id"), r.PathValue("mem_id"))
	if err != nil {
		return response.NewApiResponse(nil, err)
	}

	log.Debug("Member removed from team")
	return response.NewApiResponse(t, nil)
}
