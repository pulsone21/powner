package database

import (
	"errors"
	"fmt"

	"github.com/pulsone21/powner/internal/entities"
	"gorm.io/gorm"
)

type DBMemberRepository struct {
	db *gorm.DB
}

func (r DBMemberRepository) GetAll() (*[]entities.Member, error) {
	mems := []entities.Member{}
	res := r.db.Model(&entities.Member{}).Preload("Skills").Preload("Skills.Skill").Find(&mems)

	if res.Error != nil {
		return nil, res.Error
	}

	if res.RowsAffected == 0 {
		return &[]entities.Member{}, nil
	}

	return &mems, nil
}

func (r DBMemberRepository) GetByID(id uint) (*entities.Member, error) {
	var member entities.Member
	res := r.db.Model(&entities.Member{}).Preload("Skills").Preload("Skills.Skill").Where("Id = ?", id).First(&member)

	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if res.Error != nil {
		return nil, res.Error
	}

	return &member, nil
}

func (r DBMemberRepository) Create(newTeam entities.Member) (*entities.Member, error) {
	res := r.db.Create(&newTeam)

	if res.Error != nil {
		return nil, res.Error
	}

	if res.RowsAffected == 0 {
		return nil, fmt.Errorf("couldn't create member.")
	}

	return &newTeam, nil
}

func (r DBMemberRepository) Update(newMem entities.Member) (*entities.Member, error) {
	oldM, err := r.GetByID(newMem.ID)
	if err != nil {
		return nil, err
	}

	oldM = &newMem
	return &newMem, r.db.Save(&oldM).Error
}

func (r DBMemberRepository) Delete(id uint) error {
	return r.db.Delete(&entities.Member{}, id).Error
}

func (r DBMemberRepository) AddSkill(mem_id uint, skill entities.Skill) (*entities.Member, error) {
	m, err := r.GetByID(mem_id)
	if err != nil {
		return nil, err
	}
	err = r.db.Model(&m).Association("Skills").Append(skill)
	if err != nil {
		return nil, err
	}

	m.Skills = append(m.Skills, *entities.NewSkillRating(mem_id, skill))
	return m, nil
}

func (r DBMemberRepository) UpdateSkillRating(skillrating_id uint, rating int) error {
	res := r.db.Model(entities.SkillRating{}).Preload("Skill").Where("Id = ?", skillrating_id).Update("rating", rating)
	return res.Error
}
