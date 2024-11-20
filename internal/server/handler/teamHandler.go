package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/pulsone21/powner/internal/entities"
	"github.com/pulsone21/powner/internal/server/middleware"
	"github.com/pulsone21/powner/internal/server/response"
	"github.com/pulsone21/powner/internal/service"
)

type TeamHandler struct {
	service service.TeamService
}

// GetTeams Get all teams
//
//	@Summary		Get all teams
//	@Description	Gets all teams which are saved in the database
//	@Tags			Team
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}		entities.Team
//	@Failure		400	{object}	response.ApiResponse
//	@Failure		500	{object}	response.ApiResponse
//	@Router			/team [get]
func (h TeamHandler) getTeams(w http.ResponseWriter, r *http.Request) response.IResponse {
	log := middleware.GetLogger(r.Context())
	log.Debug("Get Teams called")

	mem, err := h.service.GetTeams()
	if err != nil {
		return *response.NewApiResponse(nil, err)
	}

	log.Debug("Teams queried")
	if len(*mem) == 0 {
		return *response.NewApiResponse(nil, nil)
	}
	log.Debug("should have some Teams")

	return *response.NewApiResponse(mem, nil)
}

// GetTeamById Get a team by its id
//
//	@Summary		gets a team by its id
//	@Description	gets a team by its id, which is the primary key in the database
//	@Tags			Team
//	@Param			id	path	int	true	"team Id"
//	@Produce		json
//	@Success		200	{array}		entities.Team
//	@Failure		400	{object}	response.ApiResponse
//	@Failure		500	{object}	response.ApiResponse
//	@Router			/team/{id} [get]
func (h TeamHandler) getTeamById(w http.ResponseWriter, r *http.Request) response.IResponse {
	log := middleware.GetLogger(r.Context())
	log.Debug("Get teams by ID hit")

	mem, err := h.service.GetTeamByID(r.PathValue("id"))
	if err != nil {
		return *response.NewApiResponse(nil, err)
	}

	log.Debug("teams queried")
	if mem == nil {
		return *response.NewApiResponse(nil, nil)
	}

	log.Debug("Should have some teams")
	return *response.NewApiResponse(mem, nil)
}

// CreateTeam Create a team
//
//	@Summary		Create a team
//	@Description	Get all teams
//	@Tags			Team
//	@Param			teamRequest	body	entities.TeamRequest	true	"Team request"
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}		entities.Team
//	@Failure		400	{object}	response.ApiResponse
//	@Failure		500	{object}	response.ApiResponse
//	@Router			/team [post]
func (h TeamHandler) createTeam(w http.ResponseWriter, r *http.Request) response.IResponse {
	log := middleware.GetLogger(r.Context())
	log.Debug("Create team hit")

	var teamReq entities.TeamRequest
	err := json.NewDecoder(r.Body).Decode(&teamReq)
	if err != nil {
		return *response.NewApiResponse(nil, errors.Join(service.BadRequest, err))
	}

	log.Debug("team request, parsed", "teamReq: ", fmt.Sprintf("%+v", teamReq))

	mem, err := h.service.CreateTeam(teamReq)
	if err != nil {
		return *response.NewApiResponse(nil, err)
	}

	log.Debug("New team created")

	return *response.NewApiResponse(mem, nil)
}

// DeleteTeam Deletes team by its id
//
//	@Summary		Deletes team by its id
//	@Description	Deletes a team by its id, which is the primary key in the database
//	@Tags			Team
//	@Param			id	path	int	true	"team Id"
//	@Produce		json
//	@Success		200	{array}		entities.Team
//	@Failure		400	{object}	response.ApiResponse
//	@Failure		500	{object}	response.ApiResponse
//	@Router			/team/{id} [delete]
func (h TeamHandler) deleteTeam(w http.ResponseWriter, r *http.Request) response.IResponse {
	log := middleware.GetLogger(r.Context())
	log.Debug("Delete team hit")

	err := h.service.DeleteTeam(r.PathValue("id"))

	log.Debug("delete team dispatched to service")
	return *response.NewApiResponse("Done", err)
}

// UpdateTeam Updates team by its id
//
//	@Summary		Updates team by its id
//	@Description	Updates a team by its id, which is the primary key in the database
//	@Tags			Team
//	@Param			id			path	int						true	"team Id"
//	@Param			teamRequest	body	entities.TeamRequest	true	"Team request"
//	@Produce		json
//	@Success		200	{array}		entities.Team
//	@Failure		400	{object}	response.ApiResponse
//	@Failure		500	{object}	response.ApiResponse
//	@Router			/team/{id} [post]
func (h TeamHandler) updateTeam(w http.ResponseWriter, r *http.Request) response.IResponse {
	log := middleware.GetLogger(r.Context())
	log.Debug("Update team hit")

	var teamReq entities.TeamRequest
	err := json.NewDecoder(r.Body).Decode(&teamReq)
	if err != nil {
		return *response.NewApiResponse(nil, errors.Join(service.BadRequest, err))
	}

	log.Debug("team request, parsed", "teamReq: ", fmt.Sprintf("%+v", teamReq))

	newTeam, err := h.service.UpdateTeam(r.PathValue("id"), teamReq)

	log.Debug("team update dispatched")
	return *response.NewApiResponse(newTeam, err)
}

func (h TeamHandler) RegisterRoutes(t *http.ServeMux) {
	t.HandleFunc("GET /team", setupApiHandler(h.getTeams))
	t.HandleFunc("POST /team", setupApiHandler(h.createTeam))
	t.HandleFunc("GET /team/{id}", setupApiHandler(h.getTeamById))
	t.HandleFunc("DELTE /team/{id}", setupApiHandler(h.deleteTeam))
	t.HandleFunc("POST /team/{id}", setupApiHandler(h.updateTeam))
}
