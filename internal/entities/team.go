package entities

import (
	"fmt"
	"log/slog"
	"sort"

	"github.com/pulsone21/powner/internal/errx"
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

func (t *Team) HasChanges(name, description string, skills *[]Skill, members *[]Member) (*Team, bool) {
	changes := false

	if t.Name != name {
		changes = true
		t.Name = name
	}

	if t.Description != description {
		changes = true
		t.Description = description
	}

	if skills != nil {
		skillChange := t.skillsChanged(*skills)
		if skillChange {
			changes = true
			t.Skills = *skills
		}
	}

	if members != nil {
		memChange := t.memberChanged(*members)
		if memChange {
			changes = true
			t.Members = *members
		}
	}

	return t, changes
}

func (t Team) skillsChanged(newS []Skill) bool {
	if len(newS) != len(t.Skills) {
		return true
	}

	sort.Sort(skillSort(newS))
	sort.Sort(skillSort(t.Skills))

	for i, s := range t.Skills {
		if s != newS[i] {
			return true
		}
	}
	return false
}

func (t Team) memberChanged(newM []Member) bool {
	if len(newM) != len(t.Members) {
		return true
	}

	sort.Sort(memberSort(newM))
	sort.Sort(memberSort(t.Members))

	for i, m := range t.Members {
		if m.ID != newM[i].ID {
			return true
		}
	}
	return false
}

type TeamRequest struct {
	Name        string
	Description string
	Skills      *[]Skill
	Members     *[]Member
}

func (t TeamRequest) ValidateFields() *errx.ErrorMap {
	var validationErrors errx.ErrorMap
	if len(t.Name) < 4 {
		validationErrors.Set("name", fmt.Errorf("name musst be longer then 3 characters"))
	}

	if len(t.Description) < 10 {
		validationErrors.Set("description", fmt.Errorf("description musst be longer then 10 characters"))
	}

	return &validationErrors
}
