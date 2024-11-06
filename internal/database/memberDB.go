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

func (r DBMemberRepository) GetByID(id string) (*entities.Member, error) {
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
	oldM, err := r.GetByID(fmt.Sprint(newMem.ID))
	if err != nil {
		return nil, err
	}

	oldM = &newMem
	return &newMem, r.db.Save(&oldM).Error
}

func (r DBMemberRepository) Delete(id string) error {
	return r.db.Delete(&entities.Member{}, id).Error
}
