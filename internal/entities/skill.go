package entities

import (
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

type SkillType int

const (
	Hard SkillType = 0
	Soft SkillType = 1
)

type skillSort []Skill

func (s skillSort) Len() int           { return len(s) }
func (s skillSort) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s skillSort) Less(i, j int) bool { return s[i].ID > s[i].ID }
func (s skillSort) toSkills() []Skill  { return []Skill(s) }
