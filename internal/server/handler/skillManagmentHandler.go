package handler

import (
	"net/http"
	"strconv"

	"github.com/pulsone21/powner/internal/server/middleware"
	"github.com/pulsone21/powner/internal/server/response"
	"github.com/pulsone21/powner/internal/service"
)

type SkillManagementHandler struct {
	service service.SkillManagement
}

func NewSkillManagmentHandler(ser service.SkillManagement) SkillManagementHandler {
	return SkillManagementHandler{
		service: ser,
	}
}

func (h SkillManagementHandler) RegisterRoutes(t *http.ServeMux) {
	t.HandleFunc("POST /member/{mem_id}/skill/{skill_id}", setupHandler(h.addSkillToMember))
	t.HandleFunc("POST /member/{mem_id}/skill/{skill_id}/{rating}", setupHandler(h.updateSkillToMember))
	t.HandleFunc("POST /team/{team_id}/skill/{skill_id}", setupHandler(h.addSkillToTeam))
	t.HandleFunc("DELETE /team/{team_id}/skill/{skill_id}", setupHandler(h.removeSkillFromTeam))
}

// addSkillToTeam - add a skill to a team
//
//	@Summary		Adds a skill to a team
//	@Description	Adds a skill by its id to the team by the team id
//	@Tags			Team
//	@Param			skill_id	path	int	true	"Skill Id"
//	@Param			team_id		path	int	true	"Team Id"
//	@Produce		json
//	@Success		200	{array}		entities.Team
//	@Failure		400	{object}	response.ApiResponse
//	@Failure		500	{object}	response.ApiResponse
//	@Router			/team/{team_id}/skill/{skill_id} [post]
func (h SkillManagementHandler) addSkillToTeam(w http.ResponseWriter, r *http.Request) response.IResponse {
	log := middleware.GetLogger(r.Context())
	log.Debug("Add Skill to team hit")

	t, err := h.service.AddSkillToTeam(r.PathValue("team_id"), r.PathValue("skill_id"))
	if err != nil {
		return response.NewApiResponse(nil, err)
	}

	log.Debug("Skill added to team")
	return response.NewApiResponse(t, nil)
}

// removeSkillFromTeam - removes a skill from a team
//
//	@Summary		removes a skill from a team
//	@Description	removes a skill by its id from the team by the team id
//	@Tags			Team
//	@Param			skill_id	path	int	true	"Skill Id"
//	@Param			team_id		path	int	true	"Team Id"
//	@Produce		json
//	@Success		200	{array}		entities.Team
//	@Failure		400	{object}	response.ApiResponse
//	@Failure		500	{object}	response.ApiResponse
//	@Router			/team/{team_id}/skill/{skill_id} [delete]
func (h SkillManagementHandler) removeSkillFromTeam(w http.ResponseWriter, r *http.Request) response.IResponse {
	log := middleware.GetLogger(r.Context())
	log.Debug("Remove skill from team hit")

	t, err := h.service.RemoveSkillToTeam(r.PathValue("team_id"), r.PathValue("skill_id"))
	if err != nil {
		return response.NewApiResponse(nil, err)
	}

	log.Debug("Skill removed from team")
	return response.NewApiResponse(t, nil)
}

// addSkillToMember - adds a skill to a member
//
//	@Summary		adds a skill to a member
//	@Description	adds a skill by its id to the member by the member id, you can set the rating directly with the query param rating
//	@Tags			Member
//	@Param			skill_id	path	int	true	"Skill Id"
//	@Param			member_id	path	int	true	"Member Id"
//	@Param			rating		query	int	false	"rating"
//	@Produce		json
//	@Success		200	{array}		entities.Member
//	@Failure		400	{object}	response.ApiResponse
//	@Failure		500	{object}	response.ApiResponse
//	@Router			/member/{member_id}/skill/{skill_id} [post]
func (h SkillManagementHandler) addSkillToMember(w http.ResponseWriter, r *http.Request) response.IResponse {
	log := middleware.GetLogger(r.Context())
	log.Debug("Add skill to member hit")

	s_rating := r.URL.Query().Get("rating")
	rating := 1
	var err error
	if s_rating != "" {
		log.Debug("rating query found", "rating", s_rating)
		rating, err = strconv.Atoi(s_rating)
		if err != nil {
			return response.NewApiResponse(nil, err)
		}
	}

	m, err := h.service.AddSkillToMember(r.PathValue("mem_id"), r.PathValue("skill_id"), rating)
	if err != nil {
		return response.NewApiResponse(nil, err)
	}

	log.Debug("Skill added to member")
	return response.NewApiResponse(m, nil)
}

// updateSkillToMember - updates the skillrating from a member
//
//	@Summary		updates the skillrating from a member
//	@Description	updates the skillrating by the skill by the id from the member by the member id
//	@Tags			Member
//	@Param			skill_id	path	int	true	"Skill Id"
//	@Param			member_id	path	int	true	"Member Id"
//	@Param			rating		path	int	true	"rating"
//	@Produce		json
//	@Success		200	{array}		entities.Member
//	@Failure		400	{object}	response.ApiResponse
//	@Failure		500	{object}	response.ApiResponse
//	@Router			/member/{member_id}/skill/{skill_id}/{rating} [post]
func (h SkillManagementHandler) updateSkillToMember(w http.ResponseWriter, r *http.Request) response.IResponse {
	log := middleware.GetLogger(r.Context())
	log.Debug("Update skill from Member hit")

	s_rating := r.PathValue("rating")
	rating := 1
	var err error
	if s_rating != "" {
		rating, err = strconv.Atoi(s_rating)
		if err != nil {
			return response.NewApiResponse(nil, err)
		}
	}

	m, err := h.service.AddSkillToMember(r.PathValue("mem_id"), r.PathValue("skill_id"), rating)
	if err != nil {
		return response.NewApiResponse(nil, err)
	}

	log.Debug("Skill updated to member")
	return response.NewApiResponse(m, nil)
}
