package middleware

import (
	"fmt"
	"net/http"
)

type htmxValidator struct{}

func (m htmxValidator) apply(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log := GetLogger(r.Context())

		log.Debug("Validating HTMX request header")

		ok := r.Header.Get("Hx-Request")
		if ok == "" {
			log.Info("Request is not a HTMX Request - Aborting")
			http.Error(w, fmt.Errorf("invalid request").Error(), 400)
			return
		}
			
		log.Debug("Request is a HTMX Request - Continueing")
		next.ServeHTTP(w, r)
	})
}

func HtmxReqValidator() Middleware { return htmxValidator{} }
