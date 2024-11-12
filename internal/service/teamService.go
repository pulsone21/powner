package service

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/pulsone21/powner/internal/entities"
	"github.com/pulsone21/powner/internal/errx"
)

func (s Service) ValidateTeamRequest(t entities.TeamRequest) *errx.ErrorMap {
	validationErrors := t.ValidateFields()
	ts, err := s.TeamRepo.GetAll()
	if err != nil {
		validationErrors.Set("Team", errors.Join(fmt.Errorf("could not fetch teams from repository"), err))
	}

	for _, te := range *ts {
		if te.Name == t.Name {
			validationErrors.Set("name", fmt.Errorf("name is already used."))
		}
	}

	if t.Members != nil {
		memFound := false
		for _, m := range *t.Members {
			me, err := s.MemberRepo.GetByID(m.ID)
			if err != nil {
				validationErrors.Set(fmt.Sprintf("member_%v", m.ID), errors.Join(InternalError, err))
			}
			if me == nil {
				memFound = true
			}
			if !memFound {
				validationErrors.Set(fmt.Sprintf("member_%v", m.ID), fmt.Errorf("Couldn't find member with id: %v", m.ID))
				break
			}
		}
	}

	if t.Skills != nil {
		skillFound := false
		for _, sk := range *t.Skills {
			me, err := s.SkillRepo.GetByID(sk.ID)
			if err != nil {
				validationErrors.Set(fmt.Sprintf("skill_%v", sk.ID), errors.Join(InternalError, err))
			}
			if me == nil {
				skillFound = true
			}
			if !skillFound {
				validationErrors.Set(fmt.Sprintf("skill_%v", sk.ID), fmt.Errorf("Couldn't find member with id: %v", sk.ID))
				break
			}
		}
	}

	return validationErrors
}

func (s Service) CreateTeam(request entities.TeamRequest) (*entities.Team, error) {
	validationErrors := s.ValidateTeamRequest(request)
	if validationErrors != nil {
		return nil, errors.Join(BadRequest, validationErrors)
	}

	t, err := s.TeamRepo.Create(*entities.NewTeam(request.Name, request.Description))
	// TODO: Test if we can create the team directly with skill and meber do i need to add it afterwards
	return t, err
}

func (s Service) GetTeams() (*[]entities.Team, error) {
	// IDEA: Filter based on user Role? RBAC
	return s.TeamRepo.GetAll()
}

func (s Service) GetTeamByID(id string) (*entities.Team, error) {
	var validationErrors errx.ErrorMap
	fid, err := strconv.Atoi(id)
	if err != nil {
		validationErrors.Set("id", err)
	}

	if validationErrors != nil {
		return nil, errors.Join(BadRequest, validationErrors)
	}
	t, err := s.TeamRepo.GetByID(uint(fid))
	if err != nil {
		return nil, errors.Join(InternalError, err)
	}
	return t, nil
}

func (s Service) DeleteTeam(id string) error {
	var validationErrors errx.ErrorMap
	fid, err := strconv.Atoi(id)
	if err != nil {
		validationErrors.Set("id", err)
	}

	if validationErrors != nil {
		return errors.Join(BadRequest, validationErrors)
	}
	err = s.TeamRepo.Delete(uint(fid))
	if err != nil {
		return errors.Join(InternalError, err)
	}

	return nil
}

func (s Service) UpdateTeam(id string, request entities.TeamRequest) (*entities.Team, error) {
	validationErrors := s.ValidateTeamRequest(request)

	fid, err := strconv.Atoi(id)
	if err != nil {
		validationErrors.Set("id", err)
	}

	if validationErrors != nil {
		return nil, errors.Join(BadRequest, validationErrors)
	}

	oldT, err := s.TeamRepo.GetByID(uint(fid))
	if err != nil {
		return nil, errors.Join(InternalError, err)
	}

	nT, change := oldT.HasChanges(request.Name, request.Description, request.Skills, request.Members)
	if change {
		_, err = s.TeamRepo.Update(*nT)
		return nT, errors.Join(InternalError, err)
	}
	return nil, errors.Join(BadRequest, fmt.Errorf("No changes to Team: %v found.", id))
}

func (s Service) AddMemberToTeam(team_id string, member_id string) (*entities.Team, error) {
	var validationErrors errx.ErrorMap
	tID, err := strconv.Atoi(team_id)
	if err != nil {
		validationErrors.Set("team_id", err)
	}

	mID, err := strconv.Atoi(member_id)
	if err != nil {
		validationErrors.Set("member_id", err)
	}

	if validationErrors != nil {
		return nil, errors.Join(BadRequest, validationErrors)
	}

	m, err := s.MemberRepo.GetByID(uint(mID))
	if err != nil {
		return nil, errors.Join(BadRequest, fmt.Errorf("member with id: %v not found", mID))
	}

	t, err := s.TeamRepo.AddMember(uint(tID), *m)
	if err != nil {
		return nil, errors.Join(InternalError, err)
	}

	return t, nil
}

func (s Service) AddSkillToTeam(team_id, skill_id string) (*entities.Team, error) {
	var validationErrors errx.ErrorMap
	tID, err := strconv.Atoi(team_id)
	if err != nil {
		validationErrors.Set("team_id", err)
	}

	sID, err := strconv.Atoi(skill_id)
	if err != nil {
		validationErrors.Set("skill_id", err)
	}

	if validationErrors != nil {
		return nil, errors.Join(BadRequest, validationErrors)
	}

	skill, err := s.SkillRepo.GetByID(uint(sID))
	if err != nil {
		return nil, errors.Join(BadRequest, fmt.Errorf("skill with id: %v not found", sID))
	}

	t, err := s.TeamRepo.AddSkill(uint(tID), *skill)
	if err != nil {
		return nil, errors.Join(InternalError, err)
	}

	return t, nil
}

func (s Service) RemoveMemberToTeam(team_id string, member_id string) (*entities.Team, error) {
	var validationErrors errx.ErrorMap
	tID, err := strconv.Atoi(team_id)
	if err != nil {
		validationErrors.Set("team_id", err)
	}

	mID, err := strconv.Atoi(member_id)
	if err != nil {
		validationErrors.Set("member_id", err)
	}

	if validationErrors != nil {
		return nil, errors.Join(BadRequest, validationErrors)
	}

	m, err := s.MemberRepo.GetByID(uint(mID))
	if err != nil {
		return nil, errors.Join(BadRequest, fmt.Errorf("member with id: %v not found", mID))
	}

	t, err := s.TeamRepo.RemoveMember(uint(tID), *m)
	if err != nil {
		return nil, errors.Join(InternalError, err)
	}

	return t, nil
}

func (s Service) RemoveSkillToTeam(team_id, skill_id string) (*entities.Team, error) {
	var validationErrors errx.ErrorMap
	tID, err := strconv.Atoi(team_id)
	if err != nil {
		validationErrors.Set("team_id", err)
	}

	sID, err := strconv.Atoi(skill_id)
	if err != nil {
		validationErrors.Set("skill_id", err)
	}

	if validationErrors != nil {
		return nil, errors.Join(BadRequest, validationErrors)
	}

	skill, err := s.SkillRepo.GetByID(uint(sID))
	if err != nil {
		return nil, errors.Join(BadRequest, fmt.Errorf("skill with id: %v not found", sID))
	}

	t, err := s.TeamRepo.RemoveSkill(uint(tID), *skill)
	if err != nil {
		return nil, errors.Join(InternalError, err)
	}

	return t, nil
}
