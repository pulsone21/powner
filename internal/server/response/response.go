package response

import "net/http"

type IResponse interface {
	Respond(http.ResponseWriter, *http.Request)
}

type empty struct{}

// TODO: i should implement some kind of context here to give options for RBAC and better logging based on request ID and stuff like this
type ResponseFunc func(w http.ResponseWriter, r *http.Request) IResponse
