package entities

import (
	"github.com/pulsone21/powner/internal/errx"
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

func (m *Member) HasChanges(name string, age int) (*Member, bool) {
	changes := false
	if m.Name != name {
		changes = true
		m.Name = name
	}

	if m.Age != age {
		changes = true
		m.Age = age
	}

	return m, changes
}

type memberSort []Member

func (s memberSort) Len() int           { return len(s) }
func (s memberSort) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s memberSort) Less(i, j int) bool { return s[i].ID > s[j].ID }
func (s memberSort) toMember() []Member { return []Member(s) }

type MemberRequest struct {
	Name string `json:"name" schema:"name"`
	Age  int    `json:"age" schema:"age"`
}

func (m MemberRequest) ValidateFields() errx.ErrorMap {
	var validationErr errx.ErrorMap
	if m.Age < 16 {
		validationErr.Set("age", "Age must be bigger then 16, no kids labor allowed in here...")
	}

	if len(m.Name) < 3 {
		validationErr.Set("name", "Name must be >= 3 chars")
	}
	return validationErr
}
