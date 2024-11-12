package response

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"

	"github.com/pulsone21/powner/internal/service"
)

type ApiResponse struct {
	Data       any
	Error      error
	StatusCode int
}

func (res ApiResponse) Respond(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(res.StatusCode)
	if res.Error != nil {
		w.Header().Set("Content-Type", "application/json")
		slog.Error(res.Error.Error())
		w.Write([]byte(res.Error.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if res.Data == nil {
		slog.Info("Nothing found for that request")
		json.NewEncoder(w).Encode(&empty{})
		return
	}

	json.NewEncoder(w).Encode(res.Data)
}

func NewApiResponse(data any, err error) *ApiResponse {
	code := 200

	if data == nil {
		code = 404
	}

	if err != nil {
		// INFO: checking the service layer errors to get the correct statuscode
		if errors.Is(err, service.InternalError) {
			code = 500
		} else if errors.Is(err, service.BadRequest) {
			code = 400
		}
	}

	return &ApiResponse{
		Data:       data,
		Error:      err,
		StatusCode: code,
	}
}
