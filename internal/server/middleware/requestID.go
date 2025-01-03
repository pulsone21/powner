package middleware

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"net/http"
	"time"
)

type requestIdMiddleware struct{}

func (m requestIdMiddleware) apply(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		hash := generateHash(r.Method, r.URL.Path, time.Now().String(), r.RemoteAddr)
		ctx := context.WithValue(r.Context(), "requestID", hash)

		updatedReq := r.WithContext(ctx)
		next.ServeHTTP(w, updatedReq)
	})
}

func RequestID() Middleware { return requestIdMiddleware{} }

func generateHash(values ...string) string {
	h := sha256.New()
	fmt.Fprint(h, values)
	return hex.EncodeToString(h.Sum(nil))[:10]
}
