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
		return nil, errors.Join(BadRequest, validationErrors)
	}

	m, err := s.memberRepo.GetByID(uint(mID))
	if err != nil {
		return nil, errors.Join(BadRequest, fmt.Errorf("member with id: %v not found", mID))
	}

	t, err := s.teamRepo.AddMember(uint(tID), *m)
	if err != nil {
		return nil, errors.Join(InternalError, err)
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
		return nil, errors.Join(BadRequest, validationErrors)
	}

	m, err := s.memberRepo.GetByID(uint(mID))
	if err != nil {
		return nil, errors.Join(BadRequest, fmt.Errorf("member with id: %v not found", mID))
	}

	t, err := s.teamRepo.RemoveMember(uint(tID), *m)
	if err != nil {
		return nil, errors.Join(InternalError, err)
	}

	return t, nil
}
