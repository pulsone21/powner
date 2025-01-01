package handler

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/pulsone21/powner/internal/server/response"
	"github.com/pulsone21/powner/internal/service"
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
	t.HandleFunc("GET /skills/{id}/member/{mID}", setupHandler(h.handleAddSkillToMember))
	t.HandleFunc("GET /skills/{id}/team/{tID}", setupHandler(h.handleAddSkillToTeam))
	t.HandleFunc("DELETE /skills/{id}/team/{tID}", setupHandler(h.handleRemoveSkillToTeam))
}

func (h *SkillMgmtPartialsHandler) handleAddSkillToMember(w http.ResponseWriter, r *http.Request) response.IResponse {
	id := r.PathValue("id")
	sId := r.PathValue("sID")

	_, err := h.sServ.AddSkillToMember(id, sId, 0)
	if err != nil {
		return response.NewUIResponse(nil, err)
	}

	return response.NewUIResponse(templ.NopComponent, nil)
}

func (h *SkillMgmtPartialsHandler) handleAddSkillToTeam(w http.ResponseWriter, r *http.Request) response.IResponse {
	id := r.PathValue("id")
	sId := r.PathValue("sID")

	_, err := h.sServ.AddSkillToTeam(id, sId)
	if err != nil {
		return response.NewUIResponse(nil, err)
	}

	return response.NewUIResponse(templ.NopComponent, nil)
}

func (h *SkillMgmtPartialsHandler) handleRemoveSkillToTeam(w http.ResponseWriter, r *http.Request) response.IResponse {
	id := r.PathValue("id")
	sId := r.PathValue("sID")

	_, err := h.sServ.RemoveSkillToTeam(id, sId)
	if err != nil {
		return response.NewUIResponse(nil, err)
	}

	return response.NewUIResponse(templ.NopComponent, nil)
}
