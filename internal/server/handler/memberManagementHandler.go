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

func NewMemberManagementHandler(ser service.MemberManagementService) MemberManagementHandler {
	return MemberManagementHandler{
		service: ser,
	}
}

func (h MemberManagementHandler) RegisterRoutes(t *http.ServeMux) {
	t.HandleFunc("POST /team/{team_id}/member/{mem_id}", setupHandler(h.addMemberToTeam))
	t.HandleFunc("DELETE /team/{team_id}/member/{mem_id}", setupHandler(h.removeMemberFromTeam))
}

// Add Member to Team add a new Member to Team
//
//	@Summary		Adds a member to a team
//	@Description	Updates a member by its id, which is the primary key in the database
//	@Tags			Team
//	@Param			member_id	path	int	true	"Member Id"
//	@Param			team_id		path	int	true	"Team Id"
//	@Produce		json
//	@Success		200	{array}		entities.Team
//	@Failure		400	{object}	response.ApiResponse
//	@Failure		500	{object}	response.ApiResponse
//	@Router			/team/{team_id}/member/{member_id} [post]
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

// RemoveMemberfromTeam removes a Member from a Team
//
//	@Summary		remove a member from a team
//	@Description	Removes a member by its id from the team given by the id.
//	@Tags			Team
//	@Param			member_id	path	int	true	"Member Id"
//	@Param			team_id		path	int	true	"Team Id"
//	@Produce		json
//	@Success		200	{array}		entities.Team
//	@Failure		400	{object}	response.ApiResponse
//	@Failure		500	{object}	response.ApiResponse
//	@Router			/team/{team_id}/member/{member_id} [delete]
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
