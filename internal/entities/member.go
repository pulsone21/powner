package entities

import (
	"errors"
	"fmt"

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

func GetMembers(db *gorm.DB) (*[]Member, error) {
	mems := []Member{}
	res := db.Model(&Member{}).Preload("Skills").Find(&mems)

	if res.Error != nil {
		return nil, res.Error
	}

	if res.RowsAffected == 0 {
		return &[]Member{}, nil
	}

	return &mems, nil
}

func GetMemberById(db *gorm.DB, id uint) (*Member, error) {
	var member Member
	res := db.Model(&Member{}).Preload("Skills").Where("Id = ?", id).First(&member)

	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if res.Error != nil {
		return nil, res.Error
	}

	return &member, nil
}

func CreateMember(db *gorm.DB, t Member) (*Member, error) {
	res := db.Create(&t)

	if res.Error != nil {
		return nil, res.Error
	}

	if res.RowsAffected == 0 {
		return nil, fmt.Errorf("couldn't create member.")
	}

	return &t, nil
}

func DeleteMember(db *gorm.DB, id uint) error {
	return db.Delete(&Member{}, id).Error
}

func UpdateMember(db *gorm.DB, newM Member) error {
	oldM, err := GetMemberById(db, newM.ID)
	if err != nil {
		return err
	}

	oldM = &newM

	return db.Save(&oldM).Error
}
