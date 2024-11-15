package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/pulsone21/powner/internal/entities"
	"github.com/pulsone21/powner/internal/server/response"
	"github.com/pulsone21/powner/internal/service"
)

// MemberHandler is the API Handler for the member entity.
//
// responsiblity of it is to transport the http data from the body and url params to the service layer.
type MemberHandler struct {
	service service.MemberService
}

func (h MemberHandler) getMembers(w http.ResponseWriter, r *http.Request) response.IResponse {
	mem, err := h.service.GetMembers()
	if err != nil {
		return *response.NewApiResponse(nil, err)
	}

	if len(*mem) == 0 {
		return *response.NewApiResponse(nil, nil)
	}

	return *response.NewApiResponse(mem, nil)
}

func (h MemberHandler) createMember(w http.ResponseWriter, r *http.Request) response.IResponse {
	var memReq entities.MemberRequest
	err := json.NewDecoder(r.Body).Decode(&memReq)
	if err != nil {
		return *response.NewApiResponse(nil, errors.Join(service.BadRequest, err))
	}

	mem, err := h.service.CreateMember(memReq)
	if err != nil {
		return *response.NewApiResponse(nil, err)
	}

	return *response.NewApiResponse(mem, nil)
}

func (h MemberHandler) getMemberById(w http.ResponseWriter, r *http.Request) response.IResponse {
	mem, err := h.service.GetMemberByID(r.PathValue("id"))
	if err != nil {
		return *response.NewApiResponse(nil, err)
	}

	if mem == nil {
		return *response.NewApiResponse(nil, nil)
	}

	return *response.NewApiResponse(mem, nil)
}

func (h MemberHandler) deleteMember(w http.ResponseWriter, r *http.Request) response.IResponse {
	err := h.service.DeleteMember(r.PathValue("id"))
	return *response.NewApiResponse("Done", err)
}

func (h MemberHandler) updateMember(w http.ResponseWriter, r *http.Request) response.IResponse {
	var memReq entities.MemberRequest
	err := json.NewDecoder(r.Body).Decode(&memReq)
	if err != nil {
		return *response.NewApiResponse(nil, errors.Join(service.BadRequest, err))
	}
	newMem, err := h.service.UpdateMember(r.PathValue("id"), memReq)
	return *response.NewApiResponse(newMem, err)
}

func (h MemberHandler) GetRoutes() *http.ServeMux {
	mem := http.NewServeMux()
	em := http.NewServeMux()
	em.HandleFunc("GET /", setupApiHandler(h.getMembers))
	em.HandleFunc("POST /", setupApiHandler(h.createMember))
	em.HandleFunc("GET /{id}", setupApiHandler(h.getMemberById))
	em.HandleFunc("DELTE /{id}", setupApiHandler(h.deleteMember))
	em.HandleFunc("POST /{id}", setupApiHandler(h.updateMember))
	mem.Handle("/member/", http.StripPrefix("/member", em))
	return mem
}
