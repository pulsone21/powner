package handler

import (
	"fmt"
	"net/http"

	"github.com/a-h/templ"
	"github.com/pulsone21/powner/internal/entities"
	"github.com/pulsone21/powner/internal/server/middleware"
	"github.com/pulsone21/powner/internal/server/response"
	"github.com/pulsone21/powner/internal/service"
	"github.com/pulsone21/powner/internal/ui/components"
	"github.com/pulsone21/powner/internal/ui/partials"
	"github.com/pulsone21/powner/internal/ui/subpage"
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
	t.HandleFunc("GET /members/overview", setupHandler(h.serveMemberOverview))
	t.HandleFunc("GET /members/list", setupHandler(h.serveMemberList))
	t.HandleFunc("GET /members/details/{id}", setupHandler(h.serveMemberDetails))
	t.HandleFunc("DELETE /members/{id}", setupHandler(h.deleteMemberRequest))
}

func (h *MemberPartialsHandler) serveMemberOverview(w http.ResponseWriter, r *http.Request) response.IResponse {
	log := middleware.GetLogger(r.Context())
	log.Debug("team overview partial requested")

	t, err := h.mServ.GetMembers()
	if err != nil {
		return response.NewUIResponse(nil, err)
	}

	log.Debug("found all teams")
	return response.NewUIResponse(subpage.MembersOverview(*t, nil), nil)
}

func (h *MemberPartialsHandler) serveMemberDetails(w http.ResponseWriter, r *http.Request) response.IResponse {
	log := middleware.GetLogger(r.Context())
	log.Debug("member details requested")

	id := r.PathValue("id")
	t, err := h.mServ.GetMemberByID(id)
	if err != nil {
		return response.NewUIResponse(nil, err)
	}

	if t == nil {
		log.Debug("couldn't find the requested member")
		return response.NewUIResponse(nil, fmt.Errorf("Couldn't find team with id: %v", id))
	}

	log.Debug("serving member details")
	return response.NewUIResponse(subpage.MemberDetails(*t), nil)
}

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

	return response.NewUIResponse(partials.MemberList(*mems, components.DeleteMemberButton), nil)
}
