package api

import (
	"fmt"
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
	slog.Info("returning empty response")
	return newResponse(make([]interface{}, 1), nil, 200, nil)
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
