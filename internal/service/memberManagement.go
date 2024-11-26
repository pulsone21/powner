package service

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/pulsone21/powner/internal/entities"
	"github.com/pulsone21/powner/internal/errx"
	"github.com/pulsone21/powner/internal/repos"
)

type MemberManagementService struct {
	memberRepo repos.MemberRepository
	teamRepo   repos.TeamRepository
	skillRepo  repos.SkillRepository
}

func (s MemberManagementService) AddMemberToTeam(team_id string, member_id string) (*entities.Team, error) {
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
		return nil, errors.Join(BadRequest, fmt.Errorf("Validation Errors"), validationErrors)
	}

	m, err := s.memberRepo.GetByID(uint(mID))
	if err != nil {
		return nil, errors.Join(InternalError, fmt.Errorf("DB Error"), err)
	}

	if m == nil {
		return nil, errors.Join(BadRequest, fmt.Errorf("member with id: %v not found", mID))
	}

	oldT, err := s.teamRepo.GetByID(uint(tID))
	if err != nil {
		return nil, errors.Join(InternalError, fmt.Errorf("DB Error"), err)
	}
	if oldT == nil {
		return nil, errors.Join(BadRequest, fmt.Errorf("Team with id: %v not found", tID))
	}

	t, err := s.teamRepo.AddMember(*oldT, *m)
	if err != nil {
		return nil, errors.Join(InternalError, fmt.Errorf("Error on adding Member"), err)
	}

	var skillErrors errx.ErrorMap

	for _, skill := range t.Skills {
		if !m.HasSkill(skill.ID) {
			_, err := s.memberRepo.AddSkill(*m, skill)
			if err != nil {
				skillErrors.Set(
					fmt.Sprintf("Member: %v - AddSkill", m.ID),
					errors.Join(fmt.Errorf("Error in adding Skill: %v to Member: %v", skill.ID, m.ID), err))
			}
		}
	}

	if skillErrors != nil {
		return nil, errors.Join(InternalError, skillErrors)
	}
	return t, nil
}

func (s MemberManagementService) RemoveMemberToTeam(team_id string, member_id string) (*entities.Team, error) {
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
		return nil, errors.Join(BadRequest, fmt.Errorf("Validation Errors"), validationErrors)
	}

	m, err := s.memberRepo.GetByID(uint(mID))
	if err != nil {
		return nil, errors.Join(InternalError, err)
	}

	if m == nil {
		return nil, errors.Join(BadRequest, fmt.Errorf("member with id: %v not found", mID))
	}

	oldT, err := s.teamRepo.GetByID(uint(tID))
	if err != nil {
		return nil, errors.Join(InternalError, fmt.Errorf("DB Error"), err)
	}

	if oldT == nil {
		return nil, errors.Join(BadRequest, fmt.Errorf("Team with id: %v not found", tID))
	}

	if !oldT.HasMember(m.ID) {
		return nil, errors.Join(BadRequest, fmt.Errorf("Team: %v, dosen't have the Member: %v", tID, mID))
	}

	t, err := s.teamRepo.RemoveMember(*oldT, *m)
	if err != nil {
		return nil, errors.Join(InternalError, err)
	}

	return t, nil
}
