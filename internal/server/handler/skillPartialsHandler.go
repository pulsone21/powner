package handler

import (
	"fmt"
	"net/http"

	"github.com/pulsone21/powner/internal/entities"
	"github.com/pulsone21/powner/internal/server/middleware"
	"github.com/pulsone21/powner/internal/server/response"
	"github.com/pulsone21/powner/internal/service"
	"github.com/pulsone21/powner/internal/ui/partials"
)

type SkillPartialsHandler struct {
	sServ service.SkillService
	mServ service.MemberService
	tServ service.TeamService
}

func NewSkillPartialsHandler(sServ service.SkillService, tServ service.TeamService, mServ service.MemberService) *SkillPartialsHandler {
	return &SkillPartialsHandler{
		sServ: sServ,
		tServ: tServ,
		mServ: mServ,
	}
}

func (h *SkillPartialsHandler) RegisterRoutes(t *http.ServeMux) {
	t.HandleFunc("GET /skills/list", setupHandler(h.serveSkillList))
	t.HandleFunc("GET /skills/{id}", setupHandler(h.serveSkillDetails))
	t.HandleFunc("DELETE /skills/{id}", setupHandler(h.handleDeleteSkill))
}

// Path: /partials/skills/list
// QueryParams: member, team, has (eg. ?member=1&has; ?team=2)
func (h *SkillPartialsHandler) serveSkillList(w http.ResponseWriter, r *http.Request) response.IResponse {
	log := middleware.GetLogger(r.Context())
	log.Debug("skilllist partial requested")

	skills, err := h.sServ.GetSkills()
	if err != nil {
		return response.NewUIResponse(nil, err)
	}
	log.Debug("found all skills")

	var t entities.SkillHolder

	// TODO: Could implement basic quering like by Team or Skill...
	if r.URL.Query().Has("member") {
		tID := r.URL.Query().Get("member")
		t, err = h.mServ.GetMemberByID(tID)
		if err != nil {
			return response.NewUIResponse(nil, err)
		}
		if t == nil {
			return response.NewUIResponse(nil, fmt.Errorf("Member not found with id: %v", tID))
		}
	}

	if r.URL.Query().Has("team") {
		tID := r.URL.Query().Get("team")
		t, err = h.tServ.GetTeamByID(tID)
		if err != nil {
			return response.NewUIResponse(nil, err)
		}
		if t == nil {
			return response.NewUIResponse(nil, fmt.Errorf("Team not found with id: %v", tID))
		}
	}

	in := r.URL.Query().Has("has")
	final := entities.Skills(*skills).FilterByHolder(t, in)
	if in {
		return response.NewUIResponse(partials.SkillList(final.ToSkills(), t, "No skills found"), nil)
	}
	return response.NewUIResponse(partials.SkillList(final.ToSkills(), t, "No skills found"), nil)
}

// Path: /partials/skills/{id}
func (h *SkillPartialsHandler) handleDeleteSkill(w http.ResponseWriter, r *http.Request) response.IResponse {
	return nil
}

// Path: /partials/skills/{id}
func (h *SkillPartialsHandler) serveSkillDetails(w http.ResponseWriter, r *http.Request) response.IResponse {
	return nil
}
