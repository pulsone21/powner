package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/pulsone21/powner/internal/entities"
	"github.com/pulsone21/powner/internal/server/response"
	"github.com/pulsone21/powner/internal/service"
)

// SkillHandler is the API Handler for the member entity.
//
// responsiblity of it is to transport the http data from the body and url params to the service layer.
type SkillHandler struct {
	service service.SkillService
}

func (h SkillHandler) GetSkills(w http.ResponseWriter, r *http.Request) response.IResponse {
	mem, err := h.service.GetSkills()
	if err != nil {
		return *response.NewApiResponse(nil, err)
	}

	if len(*mem) == 0 {
		return *response.NewApiResponse(nil, nil)
	}

	return *response.NewApiResponse(mem, nil)
}

func (h SkillHandler) CreateSkill(w http.ResponseWriter, r *http.Request) response.IResponse {
	var skReq entities.SkillRequest
	err := json.NewDecoder(r.Body).Decode(&skReq)
	if err != nil {
		return *response.NewApiResponse(nil, errors.Join(service.BadRequest, err))
	}

	mem, err := h.service.CreateSkill(skReq)
	if err != nil {
		return *response.NewApiResponse(nil, err)
	}

	return *response.NewApiResponse(mem, nil)
}

func (h SkillHandler) GetSkillById(w http.ResponseWriter, r *http.Request) response.IResponse {
	mem, err := h.service.GetSKillByID(r.PathValue("id"))
	if err != nil {
		return *response.NewApiResponse(nil, err)
	}

	if mem == nil {
		return *response.NewApiResponse(nil, nil)
	}

	return *response.NewApiResponse(mem, nil)
}

func (h SkillHandler) DeleteSkill(w http.ResponseWriter, r *http.Request) response.IResponse {
	err := h.service.DeleteSkill(r.PathValue("id"))
	return *response.NewApiResponse("Done", err)
}

func (h SkillHandler) UpdateSkill(w http.ResponseWriter, r *http.Request) response.IResponse {
	var skReq entities.SkillRequest
	err := json.NewDecoder(r.Body).Decode(&skReq)
	if err != nil {
		return *response.NewApiResponse(nil, errors.Join(service.BadRequest, err))
	}
	newMem, err := h.service.UpdateSkill(r.PathValue("id"), skReq)
	return *response.NewApiResponse(newMem, err)
}

func (h SkillHandler) GetRoutes() *http.ServeMux {
	mem := http.NewServeMux()
	em := http.NewServeMux()
	em.HandleFunc("GET /", setupApiHandler(h.GetSkills))
	em.HandleFunc("POST /", setupApiHandler(h.CreateSkill))
	em.HandleFunc("GET /{id}", setupApiHandler(h.GetSkillById))
	em.HandleFunc("DELTE /{id}", setupApiHandler(h.DeleteSkill))
	em.HandleFunc("POST /{id}", setupApiHandler(h.UpdateSkill))
	mem.Handle("/skill/", http.StripPrefix("/skill", em))
	return mem
}
