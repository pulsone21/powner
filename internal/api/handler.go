package api

import (
	"fmt"
	"log/slog"
	"net/http"
	"path/filepath"
	"text/template"
)

func serveHtml(w http.ResponseWriter, r *http.Request) {
	route := filepath.Clean(r.URL.Path)
	slog.Info(fmt.Sprintf("requested view route: %s", route))
	_ = filepath.Join(".", "public", route, ".html")

	tmpl := template.Must(template.ParseFiles("./public/index.html"))

	tmpl.ExecuteTemplate(w, "index.html", nil)
}
