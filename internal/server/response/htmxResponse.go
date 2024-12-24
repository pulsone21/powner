package response

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/pulsone21/powner/internal/server/middleware"
	"github.com/pulsone21/powner/internal/ui/notifications"
)

type HTMXResponse struct {
	HTML       templ.Component
	Error      error
	StatusCode int
}

func (res *HTMXResponse) Respond(w http.ResponseWriter, r *http.Request) {
	log := middleware.GetLogger(r.Context())
	if res.Error != nil {
		w.Header().Set("Content-Type", "text/html")
		log.Error(res.Error.Error())
		w.WriteHeader(res.StatusCode)
		notifications.Error(res.Error.Error(), nil).Render(r.Context(), w)
		return
	}
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(res.StatusCode)

	if res.HTML == nil {
		res.HTML = templ.NopComponent
	}
	res.HTML.Render(r.Context(), w)
}

func NewUIResponse(comp templ.Component, err error) *HTMXResponse {
	return &HTMXResponse{
		HTML:       comp,
		StatusCode: evalStatusCode(comp, err),
		Error:      err,
	}
}
