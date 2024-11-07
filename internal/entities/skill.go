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

func UpdateSkill(db *gorm.DB, newS Skill) error {
	oldS, err := GetSkillById(db, newS.ID)
	if err != nil {
		return err
	}

	oldS = &newS

	return db.Save(&oldS).Error
}
