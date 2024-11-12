package handler

import (
	"net/http"

	"github.com/pulsone21/powner/internal/service"
)

type TeamHandler struct {
	service service.Service
}

// TODO: Implement the Full Handler

func (h TeamHandler) GetRoutes() *http.ServeMux {
	mem := http.NewServeMux()
	em := http.NewServeMux()
	em.HandleFunc("GET /", setupApiHandler(h.GetMembers))
	em.HandleFunc("POST /", setupApiHandler(h.CreateMember))
	em.HandleFunc("GET /{id}", setupApiHandler(h.GetMemberById))
	em.HandleFunc("DELTE /{id}", setupApiHandler(h.DeleteMember))
	em.HandleFunc("POST /{id}", setupApiHandler(h.UpdateMember))
	mem.Handle("/team/", http.StripPrefix("/team", em))
	return mem
}
