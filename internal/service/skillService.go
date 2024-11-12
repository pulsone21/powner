package service

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/pulsone21/powner/internal/entities"
	"github.com/pulsone21/powner/internal/errx"
)

func (s Service) ValidateSkillRequest(sR entities.SkillRequest) *errx.ErrorMap {
	validationErrors := sR.ValidateFields()

	skills, err := s.SkillRepo.GetAll()
	if err != nil {
		validationErrors.Set("Skill", errors.Join(fmt.Errorf("could not fetch skills from repository"), err))
	}

	for _, sk := range *skills {
		if sk.Name == sR.Name {
			validationErrors.Set("name", fmt.Errorf("Name: %v already used", sR.Name))
			break
		}
	}

	return &validationErrors
}

func (s Service) CreateSkill(request entities.SkillRequest) (*entities.Skill, error) {
	validationErrors := s.ValidateSkillRequest(request)
	if validationErrors != nil {
		return nil, errors.Join(BadRequest, validationErrors)
	}

	sR, err := s.SkillRepo.Create(*entities.NewSkill(request.Name, request.Description, entities.SkillType(request.Type), request.Importance))
	// TODO: Test if we can create the team directly with skill and meber do i need to add it afterwards
	return sR, err
}

func (s Service) GetSkills() (*[]entities.Skill, error) {
	// IDEA: Filter based on user Role? RBAC
	return s.SkillRepo.GetAll()
}

func (s Service) GetSKillByID(id string) (*entities.Skill, error) {
	var validationErrors errx.ErrorMap
	fid, err := strconv.Atoi(id)
	if err != nil {
		validationErrors.Set("id", err)
	}

	if validationErrors != nil {
		return nil, errors.Join(BadRequest, validationErrors)
	}

	m, err := s.SkillRepo.GetByID(uint(fid))
	if err != nil {
		return nil, errors.Join(InternalError, err)
	}
	return m, nil
}

func (s Service) DeleteSkill(id string) error {
	var validationErrors errx.ErrorMap
	fid, err := strconv.Atoi(id)
	if err != nil {
		validationErrors.Set("id", err)
	}

	if validationErrors != nil {
		return errors.Join(BadRequest, validationErrors)
	}
	err = s.SkillRepo.Delete(uint(fid))
	if err != nil {
		return errors.Join(InternalError, err)
	}

	return nil
}

func (s Service) UpdateSkill(id string, req entities.SkillRequest) (*entities.Skill, error) {
	validationErrors := s.ValidateSkillRequest(req)

	fid, err := strconv.Atoi(id)
	if err != nil {
		validationErrors.Set("id", err)
	}

	if validationErrors != nil {
		return nil, errors.Join(BadRequest, validationErrors)
	}

	oldT, err := s.SkillRepo.GetByID(uint(fid))
	if err != nil {
		return nil, errors.Join(InternalError, err)
	}

	nM, change := oldT.HasChanges(req.Name, req.Description, entities.SkillType(req.Type), req.Importance)
	if change {
		_, err = s.SkillRepo.Update(*nM)
		return nM, errors.Join(InternalError, err)
	}

	return nil, errors.Join(BadRequest, fmt.Errorf("No changes to Skill: %v found.", id))
}
