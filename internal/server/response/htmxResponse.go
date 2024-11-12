package response

import (
	"log/slog"
	"net/http"

	"github.com/a-h/templ"
)

type HTMXResponse struct {
	HTML       templ.Component
	Error      error
	StatusCode int
}

func (res *HTMXResponse) Respond(w http.ResponseWriter, r *http.Request) {
	if res.Error != nil {
		w.Header().Set("Content-Type", "application/json")
		slog.Error(res.Error.Error())
		w.WriteHeader(res.StatusCode)
		w.Write([]byte(res.Error.Error()))
		return
	}
	slog.Info("Request is an HTMX Request")

	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(res.StatusCode)
	res.HTML.Render(r.Context(), w)
}
