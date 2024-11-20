package server

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"
	_ "github.com/pulsone21/powner/docs"
	"github.com/pulsone21/powner/internal/database"
	"github.com/pulsone21/powner/internal/server/handler"
	mw "github.com/pulsone21/powner/internal/server/middleware"
	"github.com/pulsone21/powner/internal/server/router"
	"github.com/pulsone21/powner/internal/service"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

var errServerCreation = errors.New("Failed to create Server")

//	@title			Powner API Documentation
//	@version		1.0
//	@description	This is the api documentation of the powner application.
//
//	@BasePath		/api/v1
func CreateServer(protocol, url, port, dbPath string) (*http.Server, error) {
	var err error
	db, err := database.CreateDB(dbPath)
	if err != nil {
		return nil, errors.Join(errServerCreation, err)
	}

	envFile, _ := godotenv.Read(".env")

	teamRepo := database.NewTeamRepo(db)
	memRepo := database.NewMemberRepo(db)
	sRepo := database.NewSkillRepo(db)

	memHandler := handler.NewMemberHandler(*service.NewMemberService(memRepo))
	skillHandler := handler.NewSkillHandler(*service.NewSkillService(sRepo))
	teamHandler := handler.NewTeamHandler(*service.NewTeamService(teamRepo))
	memMgmtHandler := handler.NewMemberManagementHandler(*service.NewMemberManagement(memRepo, teamRepo, sRepo))

	apiRouter := router.NewApiRouter(1, memHandler, skillHandler, teamHandler, memMgmtHandler)

	//_ := router.NewFrontendRouter(handler.NewUIHandler(nil))

	mux := http.NewServeMux()

	loggerMW, err := mw.Logger(mw.NewLoggerConfig(envFile))
	if err != nil {
		return nil, errors.Join(errServerCreation, err)
	}

	apiChain := mw.New(
		mw.RequestID(),
		loggerMW,
	)

	mux.Handle("/api/", http.StripPrefix("/api", apiChain.Apply(apiRouter)))
	mux.Handle("/swagger/", httpSwagger.WrapHandler)
	//	mux.Handle("", uiRouter) // UI Router has already a / prefix

	s := http.Server{
		Addr:    fmt.Sprintf("%v:%v", url, port),
		Handler: mux,
	}

	log.Println(s.Addr)

	return &s, nil
}
