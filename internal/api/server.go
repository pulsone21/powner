package api

import (
	"fmt"
	"log"
	"log/slog"
	"net/http"

	"github.com/pulsone21/powner/internal/database"
	"github.com/pulsone21/powner/internal/entities"
	"github.com/pulsone21/powner/internal/middleware"
	"gorm.io/gorm"
)

var db *gorm.DB

func CreateServer(protocol, url, port, dbPath string) (*http.Server, error) {
	var err error
	db, err = database.CreateDB(dbPath)
	if err != nil {
		return nil, err
	}

	s := http.Server{
		Addr:    fmt.Sprintf("%v:%v", url, port),
		Handler: middleware.Logging(slog.Default(), getRoutes()),
	}

	log.Println(s.Addr)

	return &s, nil
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
