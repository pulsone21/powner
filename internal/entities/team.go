package entities

import (
	"errors"
	"fmt"
	"log"
	"log/slog"

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

func GetTeams(db *gorm.DB) (*[]Team, error) {
	teams := []Team{}
	res := db.Model(&Team{}).Preload("Members").Preload("Skills").Find(&teams)

	if res.Error != nil {
		return nil, res.Error
	}

	if res.RowsAffected == 0 {
		return &[]Team{}, nil
	}

	return &teams, nil
}

func GetTeamById(db *gorm.DB, id uint) (*Team, error) {
	var team Team
	res := db.Model(&Team{}).Preload("Members").Preload("Skills").Where("Id = ?", id).First(&team)

	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if res.Error != nil {
		return nil, res.Error
	}

	return &team, nil
}

func CreateTeam(db *gorm.DB, t Team) (*Team, error) {
	res := db.Create(&t)

	if res.Error != nil {
		return nil, res.Error
	}

	if res.RowsAffected == 0 {
		return nil, fmt.Errorf("couldn't create team.")
	}

	return &t, nil
}

func DeleteTeam(db *gorm.DB, id uint) error {
	t, err := GetTeamById(db, id)
	if err != nil {
		return err
	}
	err = db.Unscoped().Model(&t).Association("Members").Unscoped().Clear()
	if err != nil {
		return err
	}
	err = db.Unscoped().Model(&t).Association("Skills").Unscoped().Clear()
	if err != nil {
		return err
	}

	return db.Delete(&Team{}, id).Error
}

func UpdateTeam(db *gorm.DB, newT Team) error {
	oldT, err := GetTeamById(db, newT.ID)
	if err != nil {
		return err
	}

	log.Println(oldT)
	oldT = &newT
	log.Println(newT)
	return db.Save(&oldT).Error
}

func AddMemberToTeam(db *gorm.DB, id, mem_id uint) error {
	t, err := GetTeamById(db, id)
	if err != nil {
		return err
	}

	m, err := GetMemberById(db, mem_id)
	if err != nil {
		return err
	}

	// TODO: After the Member was added to team, he should get all skills from the team which he don't have
	return db.Model(&t).Association("Members").Append(m)
}

func RemoveMemberFromTeam(db *gorm.DB, id, mem_id uint) error {
	t, err := GetTeamById(db, id)
	if err != nil {
		return err
	}

	m, err := GetMemberById(db, mem_id)
	if err != nil {
		return err
	}

	return db.Model(&t).Association("Members").Delete(m)
}

func AddSkillToTeam(db *gorm.DB, id, skill_id uint) error {
	t, err := GetTeamById(db, id)
	if err != nil {
		return err
	}

	s, err := GetSkillById(db, skill_id)
	if err != nil {
		return err
	}

	// TODO: After the Member was added to team, he should get all skills from the team which he don't have
	return db.Model(&t).Association("Skills").Append(s)
}

func RemoveSkillFromTeam(db *gorm.DB, id, skill_id uint) error {
	t, err := GetTeamById(db, id)
	if err != nil {
		return err
	}

	s, err := GetSkillById(db, skill_id)
	if err != nil {
		return err
	}

	return db.Model(&t).Association("Skills").Delete(s)
}

func (t Team) HasSkill(skillID uint) bool {
	for _, sR := range t.Skills {
		if sR.ID == skillID {
			slog.Info("Team: %v has the Skill: %v with id: %v\n", t.Name, sR.Name, skillID)
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
