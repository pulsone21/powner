package response

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/pulsone21/powner/internal/server/middleware"
)

type HTMXResponse struct {
	HTML       templ.Component
	Error      error
	StatusCode int
}

func (res *HTMXResponse) Respond(w http.ResponseWriter, r *http.Request) {
	log := middleware.GetLogger(r.Context())
	if res.Error != nil {
		w.Header().Set("Content-Type", "application/json")
		log.Error(res.Error.Error())
		w.WriteHeader(res.StatusCode)
		w.Write([]byte(res.Error.Error()))
		return
	}

	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(res.StatusCode)
	res.HTML.Render(r.Context(), w)
}

func NewUIResponse(comp templ.Component, err error) *HTMXResponse {
	code := 200
	return &HTMXResponse{
		HTML:       comp,
		StatusCode: code,
		Error:      err,
	}
}
