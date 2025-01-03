package handler

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/a-h/templ"
	"github.com/pulsone21/powner/internal/server/middleware"
	"github.com/pulsone21/powner/internal/server/response"
	"github.com/pulsone21/powner/internal/service"
	"github.com/pulsone21/powner/internal/ui/partials"
)

type SkillMgmtPartialsHandler struct {
	sServ service.SkillManagement
}

func NewSkillMgmtPartialsHandler(sServ service.SkillManagement) *SkillMgmtPartialsHandler {
	return &SkillMgmtPartialsHandler{
		sServ: sServ,
	}
}

func (h *SkillMgmtPartialsHandler) RegisterRoutes(t *http.ServeMux) {
	t.HandleFunc("POST /skills/{id}/member/{mID}", setupHandler(h.handleAddSkillToMember))
	t.HandleFunc("POST /skills/{id}/member/{mID}/{rating}", setupHandler(h.handleUpdateSkillRating))
	t.HandleFunc("POST /skills/{id}/team/{tID}", setupHandler(h.handleAddSkillToTeam))
	t.HandleFunc("DELETE /skills/{id}/team/{tID}", setupHandler(h.handleRemoveSkillToTeam))
}

// Path: /partials/skills/{id}/member/{mID}
func (h *SkillMgmtPartialsHandler) handleAddSkillToMember(w http.ResponseWriter, r *http.Request) response.IResponse {
	id := r.PathValue("id")
	mId := r.PathValue("mID")

	_, err := h.sServ.AddSkillToMember(mId, id, 1)
	if err != nil {
		return response.NewUIResponse(nil, err)
	}

	w.Header().Add("HX-Trigger", service.ChangeMemberEvent)
	return response.NewUIResponse(templ.NopComponent, nil)
}

func (h *SkillMgmtPartialsHandler) handleAddSkillToTeam(w http.ResponseWriter, r *http.Request) response.IResponse {
	id := r.PathValue("id")
	tId := r.PathValue("tID")

	_, err := h.sServ.AddSkillToTeam(tId, id)
	if err != nil {
		return response.NewUIResponse(nil, err)
	}

	return response.NewUIResponse(templ.NopComponent, nil)
}

func (h *SkillMgmtPartialsHandler) handleRemoveSkillToTeam(w http.ResponseWriter, r *http.Request) response.IResponse {
	id := r.PathValue("id")
	tId := r.PathValue("tID")

	_, err := h.sServ.RemoveSkillToTeam(tId, id)
	if err != nil {
		return response.NewUIResponse(nil, err)
	}

	return response.NewUIResponse(templ.NopComponent, nil)
}

// Path: /partials/skills/{id}/member/{mID}/{rating}
func (h *SkillMgmtPartialsHandler) handleUpdateSkillRating(w http.ResponseWriter, r *http.Request) response.IResponse {
	log := middleware.GetLogger(r.Context())
	log.Debug("Updating SkillRating")
	id := r.PathValue("id")
	intId, err := strconv.Atoi(id)
	if err != nil {
		return response.NewUIResponse(nil, errors.Join(fmt.Errorf("Parsing skill id failed"), err))
	}

	log.Debug("skill id parsed")
	mId := r.PathValue("mID")

	rating := r.PathValue("rating")
	intRating, err := strconv.Atoi(rating)
	if err != nil {
		return response.NewUIResponse(nil, errors.Join(fmt.Errorf("Parsing Rating failed"), err))
	}

	log.Debug("rating id parsed")

	m, sErr := h.sServ.UpdateSkillRating(mId, id, intRating)
	if sErr != nil {
		log.Error("error in Updating SkillRating", "Error:", sErr.Error())
		return response.NewUIResponse(nil, sErr)
	}

	log.Debug("rating updated")
	w.Header().Add("HX-Trigger", service.ChangeMemberEvent)

	sR := m.GetSkillRatingBySkill(uint(intId))
	if sR == nil {
		return response.NewUIResponse(nil, fmt.Errorf("Couldn't get skillrating, but it was updated, please reload the side"))
	}

	log.Debug("got rating from new member -> returning")
	return response.NewUIResponse(partials.SkillAddjustItem(fmt.Sprint(m.ID), *sR), nil)
}
