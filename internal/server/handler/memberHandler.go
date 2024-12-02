package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/pulsone21/powner/internal/entities"
	"github.com/pulsone21/powner/internal/server/middleware"
	"github.com/pulsone21/powner/internal/server/response"
	"github.com/pulsone21/powner/internal/service"
)

// MemberHandler is the API Handler for the member entity.
//
// responsiblity of it is to transport the http data from the body and url params to the service layer.
type MemberHandler struct {
	service service.MemberService
}

func NewMemberHandler(ser service.MemberService) MemberHandler {
	return MemberHandler{
		service: ser,
	}
}

// GetMembers Get all Members
//
//	@Summary		Get all Members
//	@Description	Gets all members which are saved in the database
//	@Tags			Member
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}		entities.Member
//	@Failure		400	{object}	response.ApiResponse
//	@Failure		500	{object}	response.ApiResponse
//	@Router			/member [get]
func (h MemberHandler) getMembers(w http.ResponseWriter, r *http.Request) response.IResponse {
	log := middleware.GetLogger(r.Context())
	log.Info("Get Members called")

	mem, err := h.service.GetMembers()
	if err != nil {
		return *response.NewApiResponse(nil, err)
	}

	if len(*mem) == 0 {
		return *response.NewApiResponse(nil, nil)
	}

	return *response.NewApiResponse(mem, nil)
}

// CreateMember Create a Member
//
//	@Summary		Create a Member
//	@Description	Get all Members
//	@Tags			Member
//	@Param			MemberRequest	body	entities.MemberRequest	true	"Member request"
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}		entities.Member
//	@Failure		400	{object}	response.ApiResponse
//	@Failure		500	{object}	response.ApiResponse
//	@Router			/member [post]
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

// GetMemberById Get a member by its id
//
//	@Summary		gets a member by its id
//	@Description	gets a member by its id, which is the primary key in the database
//	@Tags			Member
//	@Param			id	path	int	true	"Member Id"
//	@Produce		json
//	@Success		200	{array}		entities.Member
//	@Failure		400	{object}	response.ApiResponse
//	@Failure		500	{object}	response.ApiResponse
//	@Router			/member/{id} [get]
func (h MemberHandler) getMemberById(w http.ResponseWriter, r *http.Request) response.IResponse {
	log := middleware.GetLogger(r.Context())
	log.Info("Get Member by ID hit")
	mem, err := h.service.GetMemberByID(r.PathValue("id"))
	if err != nil {
		return *response.NewApiResponse(nil, err)
	}

	log.Info("Members queried")
	if mem == nil {
		return *response.NewApiResponse(nil, nil)
	}

	log.Info("Should have a empty list")
	return *response.NewApiResponse(mem, nil)
}

// DeleteMember Deletes member by its id
//
//	@Summary		Deletes member by its id
//	@Description	Deletes a member by its id, which is the primary key in the database
//	@Tags			Member
//	@Param			id	path	int	true	"Member Id"
//	@Produce		json
//	@Success		200	{array}		entities.Member
//	@Failure		400	{object}	response.ApiResponse
//	@Failure		500	{object}	response.ApiResponse
//	@Router			/member/{id} [delete]
func (h MemberHandler) deleteMember(w http.ResponseWriter, r *http.Request) response.IResponse {
	err := h.service.DeleteMember(r.PathValue("id"))
	return *response.NewApiResponse("Done", err)
}

// UpdateMember Updates member by its id
//
//	@Summary		Updates member by its id
//	@Description	Updates a member by its id, which is the primary key in the database
//	@Tags			Member
//	@Param			id				path	int						true	"Member Id"
//	@Param			MemberRequest	body	entities.MemberRequest	true	"Member request"
//	@Produce		json
//	@Success		200	{array}		entities.Member
//	@Failure		400	{object}	response.ApiResponse
//	@Failure		500	{object}	response.ApiResponse
//	@Router			/member/{id} [post]
func (h MemberHandler) updateMember(w http.ResponseWriter, r *http.Request) response.IResponse {
	var memReq entities.MemberRequest
	err := json.NewDecoder(r.Body).Decode(&memReq)
	if err != nil {
		return *response.NewApiResponse(nil, errors.Join(service.BadRequest, err))
	}
	newMem, err := h.service.UpdateMember(r.PathValue("id"), memReq)
	return *response.NewApiResponse(newMem, err)
}

func (h MemberHandler) RegisterRoutes(t *http.ServeMux) {
	t.HandleFunc("GET /member", setupHandler(h.getMembers))
	t.HandleFunc("POST /member", setupHandler(h.createMember))
	t.HandleFunc("GET /member/{id}", setupHandler(h.getMemberById))
	t.HandleFunc("DELTE /member/{id}", setupHandler(h.deleteMember))
	t.HandleFunc("POST /member/{id}", setupHandler(h.updateMember))
}
