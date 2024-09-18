package entities

import (
	"fmt"
	"log"

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

func GetSkills(db *gorm.DB) (*[]Skill, error) {
	skills := []Skill{}
	res := db.Find(&skills)

	if res.Error != nil {
		log.Printf("Error in db query %v\n", res.Error.Error())
		return nil, res.Error
	}

	if res.RowsAffected == 0 {
		log.Println("Nothing found")
		return &[]Skill{}, nil
	}

	//rows, err := res.Rows()
	//if err != nil {
	//	log.Printf("Error in loading rows %v\n", res.Error.Error())
	//	return nil, res.Error
	//}

	//	defer rows.Close()
	//var skills []Skill
	//
	//for rows.Next() {
	//	var skill Skill
	//	err = rows.Scan(&skill)
	//	if err != nil {
	//		return nil, err
	//	}
	//
	//	skills = append(skills, skill)
	//}
	return &skills, nil
}

func GetSkillById(db *gorm.DB, id uint) (*Skill, error) {
	var Skill Skill
	res := db.Where("Id = ?", id).First(&Skill)

	if res.Error != nil {
		return nil, res.Error
	}

	if res.RowsAffected == 0 {
		return nil, nil
	}

	return &Skill, nil
}

func CreateSkill(db *gorm.DB, t Skill) (*Skill, error) {
	res := db.Create(&t)

	if res.Error != nil {
		return nil, res.Error
	}

	if res.RowsAffected == 0 {
		return nil, fmt.Errorf("couldn't create Skill.")
	}

	return &t, nil
}

func DeleteSkill(db *gorm.DB, id uint) error {
	return db.Delete(&Skill{}, id).Error
}

func UpdateSkill(db *gorm.DB, newS Skill) error {
	oldS, err := GetSkillById(db, newS.ID)
	if err != nil {
		return err
	}

	oldS = &newS

	return db.Save(&oldS).Error
}
