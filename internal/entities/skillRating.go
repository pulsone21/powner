package entities

import (
	"errors"

	"gorm.io/gorm"
)

type SkillRating struct {
	gorm.Model
	SkillID  int   `json:"skill_id"`
	Skill    Skill `json:"skill"`
	MemberID uint  `json:"member_id"`
	Rating   int   `json:"rating"`
}

func UpdateSkillRating(db *gorm.DB, ratId uint, rating int) error {
	res := db.Model(SkillRating{}).Preload("Skill").Where("Id = ?", ratId).Update("rating", rating)
	return res.Error
}

func GetSkillRating(db *gorm.DB, id uint) (*SkillRating, error) {
	var rat SkillRating
	res := db.Model(SkillRating{}).Preload("Skill").First(&rat, "Id = ?", id)

	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if res.Error != nil {
		return nil, res.Error
	}

	return &rat, nil
}

func DeleteSkillRating(db *gorm.DB, id uint) error {
	return db.Delete(&SkillRating{}, id).Error
}
