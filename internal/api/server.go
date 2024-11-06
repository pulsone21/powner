package api

import (
	"encoding/json"
	"fmt"
	"log"
	"log/slog"
	"net/http"

	"github.com/gorilla/schema"
	"github.com/pulsone21/powner/internal/database"
	"github.com/pulsone21/powner/internal/middleware"
	"gorm.io/gorm"
)

var (
	decoder = schema.NewDecoder()
	db      *gorm.DB
)

func decodeRequest[T any](r *http.Request) (*T, error) {
	content_type := r.Header.Get("content-type")
	var ret T

	if content_type == "application/x-www-form-urlencoded" {
		err := r.ParseForm()
		if err != nil {
			return nil, err
		}
		err = decoder.Decode(&ret, r.PostForm)
		return &ret, err

	} else if content_type == "application/json" {
		err := json.NewDecoder(r.Body).Decode(&ret)
		return &ret, err
	}
	return nil, fmt.Errorf("content-type: %v not available for decoding", content_type)
}

func CreateServer(protocol, url, port, dbPath string) (*http.Server, error) {
	var err error
	db, err = database.CreateDB(dbPath)
	if err != nil {
		return nil, err
	}

	s := http.Server{
		Addr: fmt.Sprintf("%v:%v", url, port),
		// TODO: Rewrite the middleware Stack, should look like this Middleware(log, auth...., getRoutes())
		Handler: middleware.Logging(slog.Default(), getRoutes()),
	}

	log.Println(s.Addr)

	return &s, nil
}
