package server

import (
	"fmt"
	"log"
	"log/slog"
	"net/http"

	"github.com/pulsone21/powner/internal/database"
	"github.com/pulsone21/powner/internal/middleware"
	"github.com/pulsone21/powner/internal/server/handler"
	"github.com/pulsone21/powner/internal/server/router"
	"github.com/pulsone21/powner/internal/service"
)

func CreateServer(protocol, url, port, dbPath string) (*http.Server, error) {
	var err error
	db, err := database.CreateDB(dbPath)
	if err != nil {
		return nil, err
	}

	teamRepo := database.NewTeamRepo(db)
	memRepo := database.NewMemberRepo(db)
	sRepo := database.NewSkillRepo(db)

	service := service.NewService(teamRepo, memRepo, sRepo)
	memHandler := handler.NewMemberHandler(*service)
	skillHandler := handler.NewSkillHandler(*service)
	teamHandler := handler.NewTeamHandler(*service)

	apiRouter := router.NewApiRouter(1, memHandler, skillHandler, teamHandler)

	uiRouter := router.NewFrontendRouter(apiRouter, handler.NewUIHandler(*service))

	s := http.Server{
		Addr: fmt.Sprintf("%v:%v", url, port),
		// TODO: Rewrite the middleware Stack, should look like this Middleware(log, auth...., getRoutes())
		Handler: middleware.Logging(slog.Default(), uiRouter),
	}

	log.Println(s.Addr)

	return &s, nil
}
