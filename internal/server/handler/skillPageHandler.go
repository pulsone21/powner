package handler

import (
	"net/http"

	"github.com/pulsone21/powner/internal/server/middleware"
	"github.com/pulsone21/powner/internal/server/response"
	"github.com/pulsone21/powner/internal/service"
	"github.com/pulsone21/powner/internal/ui/pages"
	"github.com/pulsone21/powner/internal/ui/subpage"
)

type SkillPageHandler struct {
	sServ service.SkillService
}

func NewSkillPageHandler(ser service.SkillService) SkillPageHandler {
	return SkillPageHandler{
		sServ: ser,
	}
}

func (h SkillPageHandler) GetRoutes() *http.ServeMux {
	t := http.NewServeMux()
	t.HandleFunc("GET /", setupHandler(h.generalSkillPage))
	t.HandleFunc("GET /{id}", setupHandler(h.specificSkillPage))
	return t
}

func (h *SkillPageHandler) generalSkillPage(w http.ResponseWriter, r *http.Request) response.IResponse {
	log := middleware.GetLogger(r.Context())
	log.Debug("general skill page requested")

	skills, err := h.sServ.GetSkills()
	if err != nil {
		return response.NewUIResponse(nil, err)
	}

	if ok := r.Header.Get("Hx-Request"); ok != "" {
		return response.NewUIResponse(subpage.SkillOverview(*skills, true), nil)
	}

	return response.NewUIResponse(pages.SkillOverviewPage(*skills), nil)
}

func (h *SkillPageHandler) specificSkillPage(w http.ResponseWriter, r *http.Request) response.IResponse {
	log := middleware.GetLogger(r.Context())
	log.Debug("specific skill page requested")

	s, err := h.sServ.GetSkillByID(r.PathValue("id"))
	if err != nil {
		return response.NewUIResponse(nil, err)
	}

	if ok := r.Header.Get("Hx-Request"); ok != "" {
		return response.NewUIResponse(subpage.SkillDetails(*s, true), nil)
	}

	skills, err := h.sServ.GetSkills()
	if err != nil {
		return response.NewUIResponse(nil, err)
	}

	return response.NewUIResponse(pages.SkillDetailPage(*s, *skills), nil)
}
