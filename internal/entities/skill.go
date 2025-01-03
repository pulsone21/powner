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

type Skills []Skill

func (s Skills) Len() int           { return len(s) }
func (s Skills) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s Skills) Less(i, j int) bool { return s[i].ID > s[j].ID }
func (s Skills) ToSkills() []Skill  { return []Skill(s) }
func (sK Skills) FilterByHolder(t SkillHolder, has bool) Skills {
	var final Skills
	for _, s := range sK {
		if t.HasSkill(s.ID) == has {
			final = append(final, s)
		}
	}
	return final
}

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
		validationErr.Set("type", fmt.Sprintf("SkillType is not valid, must be between: %b-%b", Hard, Soft))
	}

	if sR.Importance > 5 && sR.Importance < 1 {
		validationErr.Set("importance", "Importance needs to be between 1 and 5")
	}

	return validationErr
}
