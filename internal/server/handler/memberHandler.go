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
	service service.Service
}

func (h MemberHandler) GetMembers(w http.ResponseWriter, r *http.Request) response.IResponse {
	mem, err := h.service.GetMembers()
	if err != nil {
		return *response.NewApiResponse(nil, err)
	}

	if len(*mem) == 0 {
		return *response.NewApiResponse(nil, nil)
	}

	return *response.NewApiResponse(mem, nil)
}

func (h MemberHandler) CreateMember(w http.ResponseWriter, r *http.Request) response.IResponse {
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

func (h MemberHandler) GetMemberById(w http.ResponseWriter, r *http.Request) response.IResponse {
	mem, err := h.service.GetMemberByID(r.PathValue("id"))
	if err != nil {
		return *response.NewApiResponse(nil, err)
	}

	if mem == nil {
		return *response.NewApiResponse(nil, nil)
	}

	return *response.NewApiResponse(mem, nil)
}

func (h MemberHandler) DeleteMember(w http.ResponseWriter, r *http.Request) response.IResponse {
	err := h.service.DeleteMember(r.PathValue("id"))
	return *response.NewApiResponse("Done", err)
}

func (h MemberHandler) UpdateMember(w http.ResponseWriter, r *http.Request) response.IResponse {
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
	em.HandleFunc("GET /", setupApiHandler(h.GetMembers))
	em.HandleFunc("POST /", setupApiHandler(h.CreateMember))
	em.HandleFunc("GET /{id}", setupApiHandler(h.GetMemberById))
	em.HandleFunc("DELTE /{id}", setupApiHandler(h.DeleteMember))
	em.HandleFunc("POST /{id}", setupApiHandler(h.UpdateMember))
	mem.Handle("/member/", http.StripPrefix("/member", em))
	return mem
}
