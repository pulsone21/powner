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

	res := r.db.Model(&entities.Team{}).Preload("Members").Preload("Skills").Find(&teams)

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
	res := r.db.Model(&entities.Team{}).Preload("Members").Preload("Skills").Where("Id = ?", id).First(&team)

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
	oldT, err := r.GetByID(newTeam.ID)
	if err != nil {
		return nil, err
	}

	oldT = &newTeam
	return &newTeam, r.db.Save(&oldT).Error
}

func (r DBTeamRepository) Delete(id uint) error {
	t, err := r.GetByID(id)
	if err != nil {
		return err
	}
	err = r.db.Unscoped().Model(&t).Association("Members").Unscoped().Clear()
	if err != nil {
		return err
	}
	err = r.db.Unscoped().Model(&t).Association("Skills").Unscoped().Clear()
	if err != nil {
		return err
	}

	return r.db.Delete(&entities.Team{}, id).Error
}

func (r DBTeamRepository) RemoveMember(team_id uint, mem entities.Member) error {
	t, err := r.GetByID(team_id)
	if err != nil {
		return err
	}

	return r.db.Model(&t).Association("Members").Delete(mem)
}

func (r DBTeamRepository) AddMember(team_id uint, mem entities.Member) error {
	t, err := r.GetByID(team_id)
	if err != nil {
		return err
	}

	// TODO: After the Member was added to team, he should get all skills from the team which he don't have
	return r.db.Model(&t).Association("Members").Append(mem)
}

func (r DBTeamRepository) AddSkill(team_id uint, skill entities.Skill) error {
	t, err := r.GetByID(team_id)
	if err != nil {
		return err
	}

	// TODO: After the Member was added to team, he should get all skills from the team which he don't have
	return r.db.Model(&t).Association("Skills").Append(skill)
}

func (r DBTeamRepository) RemoveSkill(team_id uint, skill entities.Skill) error {
	t, err := r.GetByID(team_id)
	if err != nil {
		return err
	}

	return r.db.Model(&t).Association("Skills").Delete(skill)
}
