package handler

import (
	"net/http"

	"github.com/pulsone21/powner/internal/server/middleware"
	"github.com/pulsone21/powner/internal/server/response"
	"github.com/pulsone21/powner/internal/service"
	"github.com/pulsone21/powner/internal/ui/pages"
	"github.com/pulsone21/powner/internal/ui/subpage"
)

type MemberPageHandler struct {
	memService service.MemberService
}

func NewMemberPageHandler(ser service.MemberService) MemberPageHandler {
	return MemberPageHandler{
		memService: ser,
	}
}

func (h MemberPageHandler) GetRoutes() *http.ServeMux {
	t := http.NewServeMux()
	t.HandleFunc("GET /", setupHandler(h.generalMemberPage))
	t.HandleFunc("GET /{id}", setupHandler(h.specificMemberPage))
	return t
}

func (h *MemberPageHandler) generalMemberPage(w http.ResponseWriter, r *http.Request) response.IResponse {
	log := middleware.GetLogger(r.Context())
	log.Debug("general team page requested")
	mems, err := h.memService.GetMembers()
	if err != nil {
		return response.NewUIResponse(nil, err)
	}

	if ok := r.Header.Get("Hx-Request"); ok != "" {
		return response.NewUIResponse(subpage.MembersOverview(*mems, true), nil)
	}

	return response.NewUIResponse(pages.MembersOverviewPage(*mems), nil)
}

func (h *MemberPageHandler) specificMemberPage(w http.ResponseWriter, r *http.Request) response.IResponse {
	log := middleware.GetLogger(r.Context())
	log.Debug("specific team page requested")

	m, err := h.memService.GetMemberByID(r.PathValue("id"))
	if err != nil {
		return response.NewUIResponse(nil, err)
	}

	if ok := r.Header.Get("Hx-Request"); ok != "" {
		return response.NewUIResponse(subpage.MemberDetails(*m, true), nil)
	}

	return response.NewUIResponse(pages.MemberDetailsPage(*m), nil)
}
