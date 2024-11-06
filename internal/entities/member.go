package entities

import (
	"fmt"
	"log/slog"

	"gorm.io/gorm"
)

type Member struct {
	gorm.Model
	Name   string        `json:"name"`
	Age    int           `json:"age"`
	Skills []SkillRating `json:"skills"`
}

func NewMember(name string, age int) *Member {
	return &Member{
		Name:   name,
		Age:    age,
		Skills: []SkillRating{},
	}
}

func (m Member) HasSkill(skillID uint) bool {
	for _, sR := range m.Skills {
		if sR.Skill.ID == skillID {
			slog.Info(fmt.Sprintf("Mem: %v has the Skill: %v with id: %v\n", m.Name, sR.Skill.Name, skillID))
			return true
		}
	}
	return false
}

func (m Member) GetType() string {
	return "member"
}

func (m Member) GetID() uint {
	return m.ID
}

func (m Member) GetSkillRatingBySkill(id uint) *SkillRating {
	for _, sR := range m.Skills {
		if sR.Skill.ID == id {
			return &sR
		}
	}
	return nil
}
