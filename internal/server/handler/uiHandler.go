package handler

import (
	"net/http"

	"github.com/pulsone21/powner/internal/service"
)

type UIHandler struct {
	service service.Service
}

// TODO: Implement the UI Routes correctly

func (h UIHandler) GetRoutes() *http.ServeMux {
	return nil
}
