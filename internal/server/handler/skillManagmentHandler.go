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

func (h SkillManagementHandler) RegisterRoutes(t *http.ServeMux) {
	t.HandleFunc("POST /member/{mem_id}/skill/{skill_id}", setupApiHandler(h.addSkillToMember))
	t.HandleFunc("POST /member/{mem_id}/skill/{skill_id}/{rating}", setupApiHandler(h.updateSkillToMember))
	t.HandleFunc("POST /team/{team_id}/skill/{skill_id}", setupApiHandler(h.addSkillToTeam))
	t.HandleFunc("DELETE /team/{team_id}/skill/{skill_id}", setupApiHandler(h.removeSkillFromTeam))
}

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
