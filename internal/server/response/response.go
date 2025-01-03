package response

import (
	"errors"
	"net/http"

	"github.com/pulsone21/powner/internal/service"
)

type IResponse interface {
	Respond(http.ResponseWriter, *http.Request)
}

type empty struct{}

type ResponseFunc func(w http.ResponseWriter, r *http.Request) IResponse

func evalStatusCode(data any, err error) int {
	if err != nil {
		// INFO: checking the service layer errors to get the correct statuscode
		if errors.Is(err, service.InternalError) {
			return 500
		} else if errors.Is(err, service.BadRequest) {
			return 400
		}
	}
	if data == nil {
		return 202
	}
	return 200
}
