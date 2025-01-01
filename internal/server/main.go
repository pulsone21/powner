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
	"github.com/pulsone21/powner/internal/ui/pages"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

var errServerCreation = errors.New("Failed to create Server")

// @title			Powner API Documentation
// @version		1.0
// @description	This is the api documentation of the powner application.
//
// @BasePath		/api/v1
func CreateServer(protocol, url, port, dbPath string) (*http.Server, error) {
	var err error
	db, err := database.CreateDB(dbPath, nil)
	if err != nil {
		return nil, errors.Join(errServerCreation, err)
	}

	envFile, _ := godotenv.Read(".env")

	loggerMW, err := mw.Logger(mw.NewLoggerConfig(envFile))
	if err != nil {
		return nil, errors.Join(errServerCreation, err)
	}

	teamRepo := database.NewTeamRepo(db)
	memRepo := database.NewMemberRepo(db)
	sRepo := database.NewSkillRepo(db)

	apiChain := mw.New(
		mw.RequestID(),
		loggerMW,
	)

	uiChain := mw.New(
		mw.RequestID(),
		loggerMW,
	)

	htmxChain := mw.New(
		mw.RequestID(),
		loggerMW,
		mw.HtmxReqValidator(),
	)

	tServ := *service.NewTeamService(teamRepo)
	sServ := *service.NewSkillService(sRepo)
	mServ := *service.NewMemberService(memRepo)
	mgServ := *service.NewMemberManagement(memRepo, teamRepo, sRepo)
	sgServ := *service.NewSkillManagement(memRepo, teamRepo, sRepo)

	memHandler := handler.NewMemberHandler(mServ)
	skillHandler := handler.NewSkillHandler(sServ)
	teamHandler := handler.NewTeamHandler(tServ)
	memMgmtHandler := handler.NewMemberManagementHandler(mgServ)
	skillMgmtHandler := handler.NewSkillManagmentHandler(sgServ)

	apiRouter := router.NewApiRouter(1,
		memHandler,
		skillHandler,
		teamHandler,
		memMgmtHandler,
		skillMgmtHandler)

	generalPages := handler.NewGeneralPageHandler()
	teamPages := handler.NewTeamPageHandler(tServ)
	memPages := handler.NewMemberPageHandler(mServ)

	partialsRouter := router.NewPartialsRouter(
		handler.NewTeamPartialsHandler(tServ, mgServ),
		handler.NewMemberPartialsHandler(mServ, tServ),
		handler.NewSkillPartialsHandler(sServ),
		handler.NewSkillMgmtPartialsHandler(sgServ),
		handler.NewFormsHandler(mServ, sServ, tServ),
	)

	modalRouter := router.NewModalRouter(
		handler.NewModalHandler(tServ, mServ, sServ, mgServ, sgServ),
	)

	mux := http.NewServeMux()

	mux.Handle("/api/", apiChain.Apply(http.StripPrefix("/api", apiRouter)))
	mux.Handle("/swagger/", httpSwagger.WrapHandler)
	fs := http.FileServer(http.Dir("./public/static/"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))
	mux.Handle("/partials/", htmxChain.Apply(http.StripPrefix("/partials", partialsRouter)))
	mux.Handle("/modals/", htmxChain.Apply(http.StripPrefix("/modals", modalRouter)))
	mux.Handle("/teams/", uiChain.Apply(http.StripPrefix("/teams", teamPages.GetRoutes())))
	mux.Handle("/members/", uiChain.Apply(http.StripPrefix("/members", memPages.GetRoutes())))
	mux.Handle("/skills/", uiChain.Apply(http.StripPrefix("/skills", memPages.GetRoutes())))
	mux.Handle("/", uiChain.Apply(generalPages.GetRoutes()))

	s := http.Server{
		Addr: fmt.Sprintf("%v:%v", url, port),
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if _, pattern := mux.Handler(r); pattern != "" {
				mux.ServeHTTP(w, r)
			} else {
				http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(404)
					fmt.Println("404 handler triggered")
					pages.NotFound().Render(r.Context(), w)
				}).ServeHTTP(w, r)
			}
		}),
	}

	log.Println(s.Addr)

	return &s, nil
}
