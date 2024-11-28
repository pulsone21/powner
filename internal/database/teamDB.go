package database

import (
	"errors"
	"fmt"

	"github.com/pulsone21/powner/internal/entities"
	"gorm.io/gorm"
)

type DBTeamRepository struct {
	db *gorm.DB
}

func (r DBTeamRepository) GetAll() (*[]entities.Team, error) {
	teams := []entities.Team{}

	res := r.db.Model(&entities.Team{}).Preload("Members").Preload("Members.Skills").Preload("Members.Skills.Skill").Preload("Skills").Find(&teams)

	if res.Error != nil {
		return nil, res.Error
	}

	if res.RowsAffected == 0 {
		return &[]entities.Team{}, nil
	}

	return &teams, nil
}

func (r DBTeamRepository) GetByID(id uint) (*entities.Team, error) {
	var team entities.Team
	res := r.db.Model(&entities.Team{}).Preload("Members").Preload("Members.Skills").Preload("Members.Skills.Skill").Preload("Skills").Where("Id = ?", id).First(&team)

	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if res.Error != nil {
		return nil, res.Error
	}

	return &team, nil
}

func (r DBTeamRepository) Create(newTeam entities.Team) (*entities.Team, error) {
	res := r.db.Create(&newTeam)

	if res.Error != nil {
		return nil, res.Error
	}

	if res.RowsAffected == 0 {
		return nil, fmt.Errorf("couldn't create team.")
	}

	return &newTeam, nil
}

func (r DBTeamRepository) Update(newTeam entities.Team) (*entities.Team, error) {
	s := r.db.Save(&newTeam)
	if s.Error != nil {
		fmt.Println("Error in updating team to db", s.Error.Error())
		return nil, s.Error
	}
	return &newTeam, nil
}

func (r DBTeamRepository) Delete(t entities.Team) error {
	err := r.db.Unscoped().Model(&t).Association("Members").Unscoped().Clear()
	if err != nil {
		return errors.Join(
			fmt.Errorf("Could not clear member associations"), err)
	}
	err = r.db.Unscoped().Model(&t).Association("Skills").Unscoped().Clear()
	if err != nil {
		return errors.Join(
			fmt.Errorf("Could not clear skill associations"), err)
	}

	return r.db.Delete(&entities.Team{}, t.ID).Error
}

func (r DBTeamRepository) RemoveMember(t entities.Team, mem entities.Member) (*entities.Team, error) {
	err := r.db.Model(&t).Association("Members").Delete(&mem)
	if err != nil {
		return &t, err
	}

	idx := 0
	for i, m := range t.Members {
		if m.ID == mem.ID {
			idx = i
		}
	}

	if idx == 0 {
		t.Members = []entities.Member{}
		return &t, nil
	}

	t.Members = append(t.Members[:idx], t.Members[idx+1:]...)
	return &t, nil
}

func (r DBTeamRepository) AddMember(t entities.Team, mem entities.Member) (*entities.Team, error) {
	err := r.db.Model(&t).Association("Members").Append(&mem)
	if err != nil {
		return nil, err
	}

	t.Members = append(t.Members, mem)
	return &t, nil
}

func (r DBTeamRepository) AddSkill(t entities.Team, skill entities.Skill) (*entities.Team, error) {
	err := r.db.Model(&t).Association("Skills").Append(&skill)
	if err != nil {
		return nil, err
	}

	t.Skills = append(t.Skills, skill)
	return &t, nil
}

func (r DBTeamRepository) RemoveSkill(t entities.Team, skill entities.Skill) (*entities.Team, error) {
	err := r.db.Model(&t).Association("Skills").Delete(skill)
	if err != nil {
		return &t, err
	}

	idx := 0
	for i, m := range t.Skills {
		if m.ID == skill.ID {
			idx = i
		}
	}

	if idx == 0 {
		t.Skills = []entities.Skill{}
		return &t, nil
	}

	t.Skills = append(t.Skills[:idx], t.Skills[idx+1:]...)
	return &t, nil
}
