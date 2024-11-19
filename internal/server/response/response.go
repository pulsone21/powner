package response

import (
	"net/http"
)

type IResponse interface {
	Respond(http.ResponseWriter, *http.Request)
}

type empty struct{}

type ResponseFunc func(w http.ResponseWriter, r *http.Request) IResponse
