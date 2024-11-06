package database

import (
	"errors"
	"fmt"
	"log"

	"github.com/pulsone21/powner/internal/entities"
	"gorm.io/gorm"
)

type DBSkillRepository struct {
	db *gorm.DB
}

func (r DBSkillRepository) GetAll() (*[]entities.Skill, error) {
	skills := []entities.Skill{}
	res := r.db.Find(&skills)

	if res.Error != nil {
		log.Printf("Error in db query %v\n", res.Error.Error())
		return nil, res.Error
	}

	if res.RowsAffected == 0 {
		log.Println("Nothing found")
		return &[]entities.Skill{}, nil
	}

	return &skills, nil
}

func (r DBSkillRepository) GetByID(id string) (*entities.Skill, error) {
	var skill entities.Skill
	res := r.db.Where("Id = ?", id).First(&skill)

	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if res.Error != nil {
		return nil, res.Error
	}

	return &skill, nil
}

func (r DBSkillRepository) Create(newSkill entities.Skill) (*entities.Skill, error) {
	res := r.db.Create(&newSkill)

	if res.Error != nil {
		return nil, res.Error
	}

	if res.RowsAffected == 0 {
		return nil, fmt.Errorf("couldn't create Skill.")
	}

	return &newSkill, nil
}

func (r DBSkillRepository) Update(newSkill entities.Skill) (*entities.Skill, error) {
	oldS, err := r.GetByID(fmt.Sprint(newSkill.ID))
	if err != nil {
		return nil, err
	}

	oldS = &newSkill

	return &newSkill, r.db.Save(&oldS).Error
}

func (r DBSkillRepository) Delete(id string) error {
	return r.db.Delete(&entities.Skill{}, id).Error
}
