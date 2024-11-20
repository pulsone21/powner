package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/pulsone21/powner/internal/entities"
	"github.com/pulsone21/powner/internal/server/response"
	"github.com/pulsone21/powner/internal/service"
)

// SkillHandler is the API Handler for the Skill entity.
//
// responsiblity of it is to transport the http data from the body and url params to the service layer.
type SkillHandler struct {
	service service.SkillService
}

// GetSkills Get all Skills
//
//	@Summary		Get all skill
//	@Description	Gets all kill which are saved in the database
//	@Tags			Skill
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}		entities.Skill
//	@Failure		400	{object}	response.ApiResponse
//	@Failure		500	{object}	response.ApiResponse
//	@Router			/skill [get]
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

// CreateSkill Create a Skill
//
//	@Summary		Create a Skill
//	@Description	Get all Skills
//	@Tags			Skill
//	@Param			SkillRequest	body	entities.MemberRequest	true	"Member request"
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}		entities.Skill
//	@Failure		400	{object}	response.ApiResponse
//	@Failure		500	{object}	response.ApiResponse
//	@Router			/skill [post]
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

// GetSkillById Get a skill by its id
//
//	@Summary		gets a Skill by its id
//	@Description	gets a Skill by its id, which is the primary key in the database
//	@Tags			Skill
//	@Param			id	path	int	true	"Skill Id"
//	@Produce		json
//	@Success		200	{array}		entities.Skill
//	@Failure		400	{object}	response.ApiResponse
//	@Failure		500	{object}	response.ApiResponse
//	@Router			/skill/{id} [get]
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

// DeleteSkill Deletes skill by its id
//
//	@Summary		Deletes Skill by its id
//	@Description	Deletes a Skill by its id, which is the primary key in the database
//	@Tags			Skill
//	@Param			id	path	int	true	"Skill Id"
//	@Produce		json
//	@Success		200	{array}		entities.Skill
//	@Failure		400	{object}	response.ApiResponse
//	@Failure		500	{object}	response.ApiResponse
//	@Router			/skill/{id} [delete]
func (h SkillHandler) DeleteSkill(w http.ResponseWriter, r *http.Request) response.IResponse {
	err := h.service.DeleteSkill(r.PathValue("id"))
	return *response.NewApiResponse("Done", err)
}

// UpdateSkill Updates skill by its id
//
//	@Summary		Updates Skill by its id
//	@Description	Updates a Skill by its id, which is the primary key in the database
//	@Tags			Skill
//	@Param			id				path	int						true	"Skill Id"
//	@Param			SkillRequest	body	entities.MemberRequest	true	"Member request"
//	@Produce		json
//	@Success		200	{array}		entities.Skill
//	@Failure		400	{object}	response.ApiResponse
//	@Failure		500	{object}	response.ApiResponse
//	@Router			/skill/{id} [post]
func (h SkillHandler) UpdateSkill(w http.ResponseWriter, r *http.Request) response.IResponse {
	var skReq entities.SkillRequest
	err := json.NewDecoder(r.Body).Decode(&skReq)
	if err != nil {
		return *response.NewApiResponse(nil, errors.Join(service.BadRequest, err))
	}
	newMem, err := h.service.UpdateSkill(r.PathValue("id"), skReq)
	return *response.NewApiResponse(newMem, err)
}

func (h SkillHandler) RegisterRoutes(t *http.ServeMux) {
	t.HandleFunc("GET /skill", setupApiHandler(h.GetSkills))
	t.HandleFunc("POST /skill", setupApiHandler(h.CreateSkill))
	t.HandleFunc("GET /skill/{id}", setupApiHandler(h.GetSkillById))
	t.HandleFunc("DELTE /skill/{id}", setupApiHandler(h.DeleteSkill))
	t.HandleFunc("POST /skill/{id}", setupApiHandler(h.UpdateSkill))
}
