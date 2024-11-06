package entities

import (
	"fmt"
	"log/slog"

	"gorm.io/gorm"
)

type Team struct {
	gorm.Model
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Members     []Member `gorm:"many2many:team_members;" json:"members"`
	Skills      []Skill  `gorm:"many2many:team_skills;" json:"skills"`
}

func NewTeam(name, description string) *Team {
	return &Team{
		Name:        name,
		Description: description,
		Members:     []Member{},
		Skills:      []Skill{},
	}
}

func (t Team) HasSkill(skillID uint) bool {
	for _, sR := range t.Skills {
		if sR.ID == skillID {
			slog.Info(fmt.Sprintf("Team: %v has the Skill: %v with id: %v\n", t.Name, sR.Name, fmt.Sprint(skillID)))
			return true
		}
	}
	return false
}

func (t Team) GetType() string {
	return "team"
}

func (t Team) GetID() uint {
	return t.ID
}

func (t Team) HasMember(id uint) bool {
	for _, m := range t.Members {
		if m.ID == id {
			return true
		}
	}
	return false
}
