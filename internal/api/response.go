package api

import (
	"encoding/json"
	"fmt"
	"log"
	"log/slog"
	"net/http"

	"github.com/a-h/templ"
)

type empty struct{}

type response struct {
	Data       any
	Html       templ.Component
	Error      error
	StatusCode int
}

func newResponse(data any, html templ.Component, statusCode int, e error) *response {
	return &response{
		Data:       data,
		Html:       html,
		StatusCode: statusCode,
		Error:      e,
	}
}

func emptyResp() *response {
	return newResponse(nil, nil, 200, nil)
}

func badRequest(err error) *response {
	return newResponse(nil, nil, 400, err)
}

func idNotValid(id string) *response {
	return newResponse(nil, nil, 400, fmt.Errorf("id is not an uint: %v", id))
}

func internalError(err error) *response {
	return newResponse(nil, nil, 500, err)
}

func success(data any, html templ.Component) *response {
	return newResponse(data, html, 200, nil)
}

type responseFunc func(w http.ResponseWriter, r *http.Request) *response

func (res *response) Respond(w http.ResponseWriter, r *http.Request) {
	if res.Error != nil {
		w.Header().Set("Content-Type", "application/json")
		log.Printf("ERROR %v\n", res.Error.Error())
		w.WriteHeader(res.StatusCode)
		w.Write([]byte(res.Error.Error()))
		return
	}

	if r.Header.Get("HX-Request") == "true" {

		if res.Html == nil {
			w.Header().Set("Content-Type", "application/json")
			log.Printf("ERROR %v\n", res.Error.Error())
			w.WriteHeader(404)
			w.Write([]byte(res.Error.Error()))
			return
		}
		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(res.StatusCode)
		log.Println("Request is an HTMX Request")
		res.Html.Render(r.Context(), w)
		return

	} else {

		w.WriteHeader(res.StatusCode)
		w.Header().Set("Content-Type", "application/json")
		if res.Data == nil {
			slog.Info("Nothing found for that request")
			json.NewEncoder(w).Encode(&empty{})
			return
		}

		json.NewEncoder(w).Encode(res)
	}
}
