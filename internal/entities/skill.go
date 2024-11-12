package entities

import (
	"fmt"

	"github.com/pulsone21/powner/internal/errx"
	"gorm.io/gorm"
)

type Skill struct {
	gorm.Model
	Name        string `json:"name"`
	Description string `json:"description"`
	Type        int    `json:"type"`
	Importance  int    `json:"importance"`
}

func NewSkill(name, description string, sType SkillType, importance int) *Skill {
	return &Skill{
		Name:        name,
		Description: description,
		Type:        int(sType),
		Importance:  importance,
	}
}

func (s *Skill) HasChanges(name, description string, sT SkillType, importance int) (*Skill, bool) {
	changes := false
	if s.Name != name {
		changes = true
		s.Name = name
	}

	if s.Description != description {
		changes = true
		s.Description = description
	}

	if s.Type != int(sT) {
		changes = true
		s.Type = int(sT)
	}

	if s.Importance != importance {
		changes = true
		s.Importance = importance
	}

	return s, changes
}

type SkillType int

const (
	Hard SkillType = 0
	Soft SkillType = 1
)

type skillSort []Skill

func (s skillSort) Len() int           { return len(s) }
func (s skillSort) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s skillSort) Less(i, j int) bool { return s[i].ID > s[j].ID }
func (s skillSort) toSkills() []Skill  { return []Skill(s) }

type SkillRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Type        int    `json:"type"`
	Importance  int    `json:"importance"`
}

func (sR SkillRequest) ValidateFields() errx.ErrorMap {
	var validationErr errx.ErrorMap
	if len(sR.Name) < 3 {
		validationErr.Set("name", "Name must be longer then 3 characters")
	}

	if len(sR.Description) < 10 {
		validationErr.Set("description", "desciption must be longer then 10 characters")
	}

	if sR.Type > int(Soft) {
		validationErr.Set("type", fmt.Sprintf("SkillType is not valid, must be: %v", fmt.Sprint(Hard, Soft)))
	}

	if sR.Importance > 5 && sR.Importance < 1 {
		validationErr.Set("importance", "Importance needs to be between 1 and 5")
	}

	return validationErr
}
