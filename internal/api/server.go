package api

import (
	"encoding/json"
	"fmt"
	"log"
	"log/slog"
	"net/http"

	"github.com/pulsone21/powner/internal/database"
	"github.com/pulsone21/powner/internal/entities"
	"github.com/pulsone21/powner/internal/middleware"
	"gorm.io/gorm"
)

type response interface{}

type empty struct{}

type responseError struct {
	StatusCode int
	Error      error
}

func newRespErr(statusCode int, e error) *responseError {
	return &responseError{
		StatusCode: statusCode,
		Error:      e,
	}
}

type responseFunc func(w http.ResponseWriter, r *http.Request) (any, *responseError)

var db *gorm.DB

func CreateServer(protocol, url, port, dbPath string) (*http.Server, error) {
	var err error
	db, err = database.CreateDB(dbPath)
	if err != nil {
		return nil, err
	}

	s := http.Server{
		Addr:    fmt.Sprintf("%v:%v", url, port),
		Handler: middleware.Logging(slog.Default(), generateHandler()),
	}

	log.Println(s.Addr)

	return &s, nil
}

func generateHandler() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handleIndex)

	api := http.NewServeMux()
	api.HandleFunc("GET /member", setupHandler(getMembers))
	api.HandleFunc("POST /member", setupHandler(createMember))
	api.HandleFunc("GET /member/{id}", setupHandler(getMemberById))
	api.HandleFunc("DELETE /member/{id}", setupHandler(deleteMember))
	api.HandleFunc("PUT /member/{id}", setupHandler(updateMember))
	api.HandleFunc("GET /member/{id}/skillrating", setupHandler(getSkillratingByMember))
	api.HandleFunc("POST /member/{id}/skillrating", setupHandler(addSkillrating))
	api.HandleFunc("PUT /member/{id}/skillrating/{rating_id}", setupHandler(updateSkillrating))
	api.HandleFunc("GET /member/{id}/skillrating/{rating_id}", setupHandler(getSkillrating))
	api.HandleFunc("DELETE /member/{id}/skillrating/{rating_id}", setupHandler(deleteSkillrating))

	api.HandleFunc("GET /skill", setupHandler(getSkills))
	api.HandleFunc("POST /skill", setupHandler(createSkill))
	api.HandleFunc("GET /skill/{id}", setupHandler(getSkillById))
	api.HandleFunc("DELETE /skill/{id}", setupHandler(deleteSkill))
	api.HandleFunc("PUT /skill/{id}", setupHandler(updateSkill))

	api.HandleFunc("GET /team", setupHandler(getTeams))
	api.HandleFunc("POST /team", setupHandler(createTeam))
	api.HandleFunc("GET /team/{id}", setupHandler(getTeamById))
	api.HandleFunc("DELETE /team/{id}", setupHandler(deleteTeam))
	api.HandleFunc("PUT /team/{id}", setupHandler(updateTeam))
	api.HandleFunc("POST /team/{id}/member/{mem_id}", setupHandler(addMember))
	api.HandleFunc("DELETE /team/{id}/member/{mem_id}", setupHandler(removeMember))
	api.HandleFunc("POST /team/{id}/skill/{skill_id}", setupHandler(addSkill))
	api.HandleFunc("DELETE /team/{id}/skill/{skill_id}", setupHandler(removeSkill))

	mux.Handle("/api/", http.StripPrefix("/api", api))

	return mux
}

func setupHandler(fn responseFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		res, err := fn(w, r)
		w.Header().Set("Content-Type", "application/json")

		if err != nil {
			log.Printf("ERROR %v\n", err.Error.Error())
			w.WriteHeader(err.StatusCode)
			w.Write([]byte(err.Error.Error()))
			return
		}

		w.WriteHeader(http.StatusOK)

		if res == nil {
			slog.Info("Nothing found for that request")
			json.NewEncoder(w).Encode(&empty{})
			return
		}

		json.NewEncoder(w).Encode(res)
	})
}

func GenerateData() {
	n := 0
	log.Println("Generating Data")
	for n < 5 {
		mem := entities.NewMember("Test Member", 25+n)
		entities.CreateMember(db, *mem)
	}
	log.Println("Should be finished with data generation")
}
