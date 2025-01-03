package handler

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/pulsone21/powner/internal/entities"
	"github.com/pulsone21/powner/internal/server/middleware"
	"github.com/pulsone21/powner/internal/server/response"
	"github.com/pulsone21/powner/internal/service"
	"github.com/pulsone21/powner/internal/ui/modals"
	"github.com/pulsone21/powner/internal/ui/partials"
)

type FormsHandler struct {
	mServ service.MemberService
	sServ service.SkillService
	tServ service.TeamService
}

func NewFormsHandler(mServ service.MemberService, sServ service.SkillService, tServ service.TeamService) *FormsHandler {
	return &FormsHandler{
		mServ: mServ,
		sServ: sServ,
		tServ: tServ,
	}
}

func (h *FormsHandler) RegisterRoutes(t *http.ServeMux) {
	t.HandleFunc("GET /forms/teams/add", setupHandler(h.serveCreateTeamForm))
	t.HandleFunc("GET /forms/members/add", setupHandler(h.serveCreateMemberForm))
	t.HandleFunc("GET /forms/skills/add", setupHandler(h.serveCreateSkillForm))
	t.HandleFunc("POST /forms/members/add", setupHandler(h.handleMemberFormRequest))
	t.HandleFunc("POST /forms/teams/add", setupHandler(h.handleTeamFormRequest))
	t.HandleFunc("POST /forms/skills/add", setupHandler(h.handleSkillFormRequest))
}

func (h *FormsHandler) serveCreateTeamForm(w http.ResponseWriter, r *http.Request) response.IResponse {
	return response.NewUIResponse(modals.NewTeamModal(), nil)
}

func (h *FormsHandler) serveCreateMemberForm(w http.ResponseWriter, r *http.Request) response.IResponse {
	return response.NewUIResponse(modals.NewMemberModal(), nil)
}

func (h *FormsHandler) serveCreateSkillForm(w http.ResponseWriter, r *http.Request) response.IResponse {
	return response.NewUIResponse(modals.NewSkillModal(), nil)
}

func (h *FormsHandler) handleMemberFormRequest(w http.ResponseWriter, r *http.Request) response.IResponse {
	log := middleware.GetLogger(r.Context())

	if err := r.ParseForm(); err != nil {
		return response.NewUIResponse(nil, errors.Join(fmt.Errorf("couldn't parse form\n"), err))
	}

	log.Debug("Add MemberForm Request")
	var req entities.MemberRequest

	err := decoder.Decode(&req, r.PostForm)
	if err != nil {
		log.Error("Form decoding error", "Error", err)
		return response.NewUIResponse(nil, errors.Join(fmt.Errorf("couldn't parse form\n"), err))
	}

	log.Debug("Decoded FormData to MemberRequest")
	log.Debug(fmt.Sprintf("%+v", req))

	_, sErr := h.mServ.CreateMember(req)
	if sErr != nil {
		log.Error("Couldn't create Member", "ValidationErrors:", err)
		return response.NewUIResponse(partials.MemberForm(*sErr.GetValidationErrors()), nil)
	}

	log.Debug("Created new Member based on Request")
	w.Header().Add("HX-Trigger", service.CreateMemberEvent)
	return response.NewUIResponse(partials.MemberForm(nil), nil)
}

func (h *FormsHandler) handleTeamFormRequest(w http.ResponseWriter, r *http.Request) response.IResponse {
	log := middleware.GetLogger(r.Context())

	if err := r.ParseForm(); err != nil {
		return response.NewUIResponse(nil, errors.Join(fmt.Errorf("couldn't parse form\n"), err))
	}

	log.Debug("Add TeamForm Request")
	var req entities.TeamRequest

	err := decoder.Decode(&req, r.PostForm)
	if err != nil {
		log.Error("Form decoding error", "Error", err)
		return response.NewUIResponse(nil, errors.Join(fmt.Errorf("couldn't parse form\n"), err))
	}

	log.Debug("Decoded FormData to MemberRequest")
	log.Debug(fmt.Sprintf("%+v", req))

	_, sErr := h.tServ.CreateTeam(req)
	if sErr != nil {
		log.Error("Couldn't create team", "ValidationErrors:", err)
		return response.NewUIResponse(partials.TeamForm(*sErr.GetValidationErrors()), nil)
	}

	log.Debug("Created new team based on Request")
	w.Header().Add("HX-Trigger", service.CreateTeamEvent)
	return response.NewUIResponse(partials.TeamForm(nil), nil)
}

func (h *FormsHandler) handleSkillFormRequest(w http.ResponseWriter, r *http.Request) response.IResponse {
	log := middleware.GetLogger(r.Context())

	if err := r.ParseForm(); err != nil {
		return response.NewUIResponse(nil, errors.Join(fmt.Errorf("couldn't parse form\n"), err))
	}

	log.Debug("Add SkillForm Request")
	var req entities.SkillRequest

	err := decoder.Decode(&req, r.PostForm)
	if err != nil {
		log.Error("Form decoding error", "Error", err)
		return response.NewUIResponse(nil, errors.Join(fmt.Errorf("couldn't parse form\n"), err))
	}

	log.Debug("Decoded FormData to SkillRequest")
	log.Debug(fmt.Sprintf("%+v", req))

	_, sErr := h.sServ.CreateSkill(req)
	if sErr != nil {
		log.Error("Couldn't create skill", "ValidationErrors:", err)
		return response.NewUIResponse(partials.SkillForm(*sErr.GetValidationErrors()), nil)
	}

	log.Debug("Created new skill based on Request")
	w.Header().Add("HX-Trigger", service.CreateSkillEvent)
	return response.NewUIResponse(partials.SkillForm(nil), nil)
}
