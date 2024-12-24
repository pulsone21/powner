package response

import (
	"encoding/json"
	"net/http"

	"github.com/pulsone21/powner/internal/server/middleware"
)

type ApiResponse struct {
	Data       any
	Error      error
	StatusCode int
}

func (res ApiResponse) Respond(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	log := middleware.GetLogger(r.Context())
	if res.Error != nil {
		w.WriteHeader(res.StatusCode)
		log.Error(res.Error.Error())
		w.Write([]byte(res.Error.Error()))
		return
	}

	if res.Data == nil {
		log.Info("Nothing found for that request")
		json.NewEncoder(w).Encode(&empty{})
		return
	}

	json.NewEncoder(w).Encode(res.Data)
}

func NewApiResponse(data any, err error) *ApiResponse {
	return &ApiResponse{
		Data:       data,
		Error:      err,
		StatusCode: evalStatusCode(data, err),
	}
}
