package handler

import (
	"fmt"
	"net/http"

	"github.com/pulsone21/powner/internal/entities"
	"github.com/pulsone21/powner/internal/server/middleware"
	"github.com/pulsone21/powner/internal/server/response"
	"github.com/pulsone21/powner/internal/service"
	"github.com/pulsone21/powner/internal/ui/modals"
)

type ModalHandler struct {
	tServ  service.TeamService
	mServ  service.MemberService
	sServ  service.SkillService
	mgServ service.MemberManagementService
	sgServ service.SkillManagement
}

func NewModalHandler(
	tServ service.TeamService,
	mServ service.MemberService,
	sServ service.SkillService,
	mgServ service.MemberManagementService,
	sgServ service.SkillManagement,
) *ModalHandler {
	return &ModalHandler{
		tServ:  tServ,
		mServ:  mServ,
		sServ:  sServ,
		mgServ: mgServ,
		sgServ: sgServ,
	}
}

func (h *ModalHandler) RegisterRoutes(t *http.ServeMux) {
	t.HandleFunc("GET /members", setupHandler(h.serveMemberAssignModal))
	t.HandleFunc("GET /skills", setupHandler(h.serveSkillAssignModal))
}

func (h *ModalHandler) serveMemberAssignModal(w http.ResponseWriter, r *http.Request) response.IResponse {
	log := middleware.GetLogger(r.Context())
	log.Debug("Assign Member Modal requested")

	tID := r.URL.Query().Get("team")
	t, err := h.tServ.GetTeamByID(tID)
	if err != nil {
		return response.NewUIResponse(nil, err)
	}
	log.Debug(fmt.Sprintf("Looking for member which are not in %v", t.Name))

	mems, err := h.mServ.GetMembers()
	if err != nil {
		return response.NewUIResponse(nil, err)
	}

	var finalM []entities.Member
	for _, m := range *mems {
		if !t.HasMember(m.ID) {
			finalM = append(finalM, m)
		}
	}

	log.Debug("filtered out member which are already on the team")
	fmt.Println(finalM)

	return response.NewUIResponse(modals.MemberModal(t), nil)
}

// Path: /modals/skills.
//
// Query Params: member, team (eg. member=1, team=2).
func (h *ModalHandler) serveSkillAssignModal(w http.ResponseWriter, r *http.Request) response.IResponse {
	log := middleware.GetLogger(r.Context())
	log.Debug("Assign skill modal requested")

	sk, err := h.sServ.GetSkills()
	if err != nil {
		return response.NewUIResponse(nil, err)
	}

	if r.URL.Query().Has("member") {
		log.Debug("Got a member")
		mID := r.URL.Query().Get("member")
		m, err := h.mServ.GetMemberByID(mID)
		if err != nil {
			return response.NewUIResponse(nil, err)
		}
		var finalSk []entities.Skill
		for _, s := range *sk {
			if !m.HasSkill(s.ID) {
				finalSk = append(finalSk, s)
			}
		}

		return response.NewUIResponse(modals.SkillModal(finalSk, m), nil)
	}

	if r.URL.Query().Has("team") {
		log.Debug("Got a team")
		tID := r.URL.Query().Get("team")
		t, err := h.tServ.GetTeamByID(tID)
		if err != nil {
			return response.NewUIResponse(nil, err)
		}
		var finalSk []entities.Skill
		for _, s := range *sk {
			if !t.HasSkill(s.ID) {
				finalSk = append(finalSk, s)
			}
		}

		return response.NewUIResponse(modals.SkillModal(finalSk, t), nil)
	}

	return response.NewUIResponse(nil, fmt.Errorf("Query params unknown or not present"))
}
