package api

import (
	"net/http"

	"github.com/pulsone21/powner/internal/ui/pages"
)

func indexPage(w http.ResponseWriter, r *http.Request) {
	pages.Index().Render(r.Context(), w)
}

func writeError(err error, w http.ResponseWriter) {
	w.WriteHeader(400)
	w.Write([]byte(err.Error()))
}
