package database

import (
	"github.com/pulsone21/powner/internal/entities"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func CreateDB(dbPath string, conf *gorm.Config) (*gorm.DB, error) {
	var db *gorm.DB
	var err error
	if conf == nil {
		db, err = gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	} else {
		db, err = gorm.Open(sqlite.Open(dbPath), conf)
	}
	if err != nil {
		return nil, err
	}

	migrate(db)

	return db, nil
}

func migrate(db *gorm.DB) {
	db.AutoMigrate(&entities.Member{})
	db.AutoMigrate(&entities.SkillRating{})
	db.AutoMigrate(&entities.Team{})
	db.AutoMigrate(&entities.Skill{})
}

func NewTeamRepo(db *gorm.DB) *DBTeamRepository {
	return &DBTeamRepository{
		db: db,
	}
}

func NewMemberRepo(db *gorm.DB) *DBMemberRepository {
	return &DBMemberRepository{
		db: db,
	}
}

func NewSkillRepo(db *gorm.DB) *DBSkillRepository {
	return &DBSkillRepository{
		db: db,
	}
}
