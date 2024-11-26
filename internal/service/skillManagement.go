package service

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/pulsone21/powner/internal/entities"
	"github.com/pulsone21/powner/internal/errx"
	"github.com/pulsone21/powner/internal/repos"
)

type SkillManagement struct {
	memberRepo repos.MemberRepository
	teamRepo   repos.TeamRepository
	skillRepo  repos.SkillRepository
}

func (s SkillManagement) RemoveSkillToTeam(team_id, skill_id string) (*entities.Team, error) {
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

	skill, err := s.skillRepo.GetByID(uint(sID))
	if err != nil {
		return nil, errors.Join(InternalError, fmt.Errorf("DB Error"), err)
	}

	if skill == nil {
		return nil, errors.Join(BadRequest, fmt.Errorf("skill with id: %v not found", sID))
	}

	oldT, err := s.teamRepo.GetByID(uint(tID))
	if err != nil {
		return nil, errors.Join(InternalError, fmt.Errorf("DB Error"), err)
	}
	if oldT == nil {
		return nil, errors.Join(BadRequest, fmt.Errorf("Team with id: %v not found", tID))
	}

	if !oldT.HasSkill(skill.ID) {
		return nil, errors.Join(BadRequest, fmt.Errorf("team with id: %v didn't has the skill: %v", oldT.ID, skill.ID))
	}

	t, err := s.teamRepo.RemoveSkill(*oldT, *skill)
	if err != nil {
		return nil, errors.Join(InternalError, err)
	}

	return t, nil
}

func (s SkillManagement) AddSkillToMember(mem_id, skill_id string, rating int) (*entities.Member, error) {
	var validationErrors errx.ErrorMap
	fid, err := strconv.Atoi(mem_id)
	if err != nil {
		validationErrors.Set("mem_id", err)
	}

	oldM, err := s.memberRepo.GetByID(uint(fid))
	if err != nil {
		return nil, errors.Join(InternalError, fmt.Errorf("DB Error"), err)
	}

	if oldM == nil {
		return nil, errors.Join(BadRequest, fmt.Errorf("Member with ID: %v not found", fid))
	}

	sid, err := strconv.Atoi(skill_id)
	if err != nil {
		validationErrors.Set("skill_id", err)
	}

	if validationErrors != nil {
		return nil, errors.Join(BadRequest, validationErrors)
	}

	sk, err := s.skillRepo.GetByID(uint(sid))
	if err != nil {
		return nil, errors.Join(InternalError, fmt.Errorf("DB Error"), err)
	}

	if sk == nil {
		return nil, errors.Join(BadRequest, fmt.Errorf("Skill with ID: %v not found", fid))
	}

	if oldM.HasSkill(sk.ID) {
		validationErrors.Set("skill_id", fmt.Sprintf("Member already has the skill with id: %v", skill_id))
	}

	if validationErrors != nil {
		return nil, errors.Join(BadRequest, validationErrors)
	}

	m, err := s.memberRepo.AddSkill(*oldM, *sk)
	if err != nil {
		return nil, errors.Join(InternalError, err)
	}

	if s.validRatingScore(rating) {
		m, err = s.UpdateSkillRating(fmt.Sprint(m.ID), fmt.Sprint(sk.ID), rating)
		if err != nil {
			return nil, errors.Join(InternalError, err)
		}
	}

	return m, nil
}

func (s SkillManagement) AddSkillToTeam(team_id, skill_id string) (*entities.Team, error) {
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

	skill, err := s.skillRepo.GetByID(uint(sID))
	if err != nil {
		return nil, errors.Join(InternalError, fmt.Errorf("DB Error"), err)
	}

	if skill == nil {
		return nil, errors.Join(BadRequest, fmt.Errorf("skill with id: %v not found", sID))
	}

	oldT, err := s.teamRepo.GetByID(uint(tID))
	if err != nil {
		return nil, errors.Join(InternalError, fmt.Errorf("DB Error"), err)
	}
	if oldT == nil {
		return nil, errors.Join(BadRequest, fmt.Errorf("Team with id: %v not found", tID))
	}

	t, err := s.teamRepo.AddSkill(*oldT, *skill)
	if err != nil {
		return nil, errors.Join(InternalError, err)
	}

	var memberUpdateErrs errx.ErrorMap
	for _, m := range t.Members {
		if !m.HasSkill(skill.ID) {
			_, err := s.memberRepo.AddSkill(m, *skill)
			if err != nil {
				memberUpdateErrs.Set(fmt.Sprintf("Member_%v", m.ID), err)
			}
		}
	}

	if memberUpdateErrs != nil {
		return nil, errors.Join(InternalError, memberUpdateErrs)
	}

	return s.teamRepo.GetByID(t.ID)
}

func (s SkillManagement) UpdateSkillRating(mem_id, skill_id string, rating int) (*entities.Member, error) {
	var validationErrors errx.ErrorMap
	fid, err := strconv.Atoi(mem_id)
	if err != nil {
		validationErrors.Set("mem_id", err)
	}

	if s.validRatingScore(rating) {
		validationErrors.Set("rating", fmt.Errorf("Rating needs to be between 1 - 5, is: %b", rating))
	}

	oldM, err := s.memberRepo.GetByID(uint(fid))
	if err != nil {
		return nil, errors.Join(InternalError, err)
	}

	if oldM == nil {
		return nil, errors.Join(BadRequest, fmt.Errorf("Member: %b dosn't exists", fid))
	}

	sid, err := strconv.Atoi(skill_id)
	if err != nil {
		validationErrors.Set("skill_id", err)
	}

	sRating := oldM.GetSkillRatingBySkill(uint(sid))
	if sRating == nil {
		validationErrors.Set("skill_id", fmt.Errorf("member: %b dosn't has the skill: %b", fid, sid))
	}

	if validationErrors != nil {
		return nil, errors.Join(BadRequest, validationErrors)
	}

	err = s.memberRepo.UpdateSkillRating(sRating.ID, rating)
	if err != nil {
		return nil, errors.Join(InternalError, err)
	}

	m, err := s.memberRepo.GetByID(uint(fid))
	if err != nil {
		return nil, errors.Join(InternalError, err)
	}

	return m, nil
}

func (s SkillManagement) validRatingScore(rating int) bool { return rating < 1 || rating > 5 }
