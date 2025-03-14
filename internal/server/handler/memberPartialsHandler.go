package handler

import (
	"fmt"
	"net/http"

	"github.com/a-h/templ"
	"github.com/pulsone21/powner/internal/entities"
	"github.com/pulsone21/powner/internal/server/middleware"
	"github.com/pulsone21/powner/internal/server/response"
	"github.com/pulsone21/powner/internal/service"
	"github.com/pulsone21/powner/internal/ui/partials"
)

type MemberPartialsHandler struct {
	mServ service.MemberService
	tServ service.TeamService
}

func NewMemberPartialsHandler(mServ service.MemberService, tServ service.TeamService) *MemberPartialsHandler {
	return &MemberPartialsHandler{
		mServ: mServ,
		tServ: tServ,
	}
}

func (h *MemberPartialsHandler) RegisterRoutes(t *http.ServeMux) {
	t.HandleFunc("GET /members/list", setupHandler(h.serveMemberList))
	t.HandleFunc("DELETE /members/{id}", setupHandler(h.deleteMemberRequest))
	t.HandleFunc("GET /members/{id}/skilllist", setupHandler(h.serveMemberSkillList))
}

// Path: /partials/members/{id}
func (h *MemberPartialsHandler) deleteMemberRequest(w http.ResponseWriter, r *http.Request) response.IResponse {
	log := middleware.GetLogger(r.Context())
	log.Debug("member deletion requested")
	id := r.PathValue("id")
	err := h.mServ.DeleteMember(id)
	if err != nil {
		return response.NewUIResponse(nil, err)
	}

	log.Debug("member deleted")
	w.Header().Add("HX-Trigger", service.DeleteMemberEvent)
	return response.NewUIResponse(templ.NopComponent, nil)
}

// Path: /partials/members/list
// Query Parameter: teamID, in (bool flag)
func (h *MemberPartialsHandler) serveMemberList(w http.ResponseWriter, r *http.Request) response.IResponse {
	// TODO: Refactore.... Member list has chagned a lot
	log := middleware.GetLogger(r.Context())
	log.Debug("memberlist partial requested")

	mems, err := h.mServ.GetMembers()
	if err != nil {
		return response.NewUIResponse(nil, err)
	}
	log.Debug("found all members")

	// TODO: Could implement basic quering like by Team or Skill...
	if r.URL.Query().Has("teamID") {
		tID := r.URL.Query().Get("teamID")
		t, err := h.tServ.GetTeamByID(tID)
		if err != nil {
			return response.NewUIResponse(nil, err)
		}
		if t == nil {
			return response.NewUIResponse(nil, fmt.Errorf("Team not found with id: %v", tID))
		}

		in := r.URL.Query().Has("in")
		var finalM []entities.Member
		for _, me := range *mems {
			if t.HasMember(me.ID) == in {
				finalM = append(finalM, me)
			}
		}

		if in {
			return response.NewUIResponse(partials.TeamMemberList(*t), nil)
		}
		return response.NewUIResponse(partials.MemberAddTeamList(*mems, t.ID), nil)
	}

	return response.NewUIResponse(partials.MemberList(*mems), nil)
}

// Path: /partials/members/{id}/skilllist
func (h *MemberPartialsHandler) serveMemberSkillList(w http.ResponseWriter, r *http.Request) response.IResponse {
	log := middleware.GetLogger(r.Context())
	log.Debug("member skill list requested")
	id := r.PathValue("id")
	m, err := h.mServ.GetMemberByID(id)
	if err != nil {
		return response.NewUIResponse(nil, err)
	}

	return response.NewUIResponse(partials.SkillAdjustList(*m), nil)
}
