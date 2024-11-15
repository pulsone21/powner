package server

import (
	"fmt"
	"log"
	"log/slog"
	"net/http"

	"github.com/pulsone21/powner/internal/database"
	"github.com/pulsone21/powner/internal/server/handler"
	"github.com/pulsone21/powner/internal/server/middleware"
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

	memHandler := handler.NewMemberHandler(*service.NewMemberService(memRepo))
	skillHandler := handler.NewSkillHandler(*service.NewSkillService(sRepo))
	teamHandler := handler.NewTeamHandler(*service.NewTeamService(teamRepo))

	apiRouter := router.NewApiRouter(1, memHandler, skillHandler, teamHandler)

	uiRouter := router.NewFrontendRouter(handler.NewUIHandler(nil))

	mux := http.NewServeMux()
	mux.Handle("/api/", http.StripPrefix("/api", apiRouter))
	mux.Handle("", uiRouter) // UI Router has already a / prefix

	s := http.Server{
		Addr: fmt.Sprintf("%v:%v", url, port),
		// TODO: Rewrite the middleware Stack, should look like this Middleware(log, auth...., getRoutes())
		Handler: middleware.Logging(slog.Default(), mux),
	}

	log.Println(s.Addr)

	return &s, nil
}
