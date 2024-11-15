package handler

import (
	"net/http"

	"github.com/pulsone21/powner/internal/server/response"
	"github.com/pulsone21/powner/internal/service"
)

type TeamHandler struct {
	service service.TeamService
}

func (h TeamHandler) getTeams(w http.ResponseWriter, r *http.Request) response.IResponse {
	return nil
}

func (h TeamHandler) getTeamById(w http.ResponseWriter, r *http.Request) response.IResponse {
	return nil
}

func (h TeamHandler) createTeam(w http.ResponseWriter, r *http.Request) response.IResponse {
	return nil
}

func (h TeamHandler) deleteTeam(w http.ResponseWriter, r *http.Request) response.IResponse {
	return nil
}

func (h TeamHandler) updateTeam(w http.ResponseWriter, r *http.Request) response.IResponse {
	return nil
}

func (h TeamHandler) GetRoutes() *http.ServeMux {
	mem := http.NewServeMux()
	em := http.NewServeMux()
	em.HandleFunc("GET /", setupApiHandler(h.getTeams))
	em.HandleFunc("POST /", setupApiHandler(h.createTeam))
	em.HandleFunc("GET /{id}", setupApiHandler(h.getTeamById))
	em.HandleFunc("DELTE /{id}", setupApiHandler(h.deleteTeam))
	em.HandleFunc("POST /{id}", setupApiHandler(h.updateTeam))
	mem.Handle("/team/", http.StripPrefix("/team", em))
	return mem
}
