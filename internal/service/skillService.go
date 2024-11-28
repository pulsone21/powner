package service

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/pulsone21/powner/internal/entities"
	"github.com/pulsone21/powner/internal/errx"
	"github.com/pulsone21/powner/internal/repos"
)

type SkillService struct {
	repo repos.SkillRepository
}

func (s SkillService) ValidateSkillRequest(sR entities.SkillRequest) errx.ErrorMap {
	validationErrors := sR.ValidateFields()

	skills, err := s.repo.GetAll()
	if err != nil {
		validationErrors.Set("Skill", errors.Join(fmt.Errorf("could not fetch skills from repository"), err))
	}

	for _, sk := range *skills {
		if strings.EqualFold(sk.Name, sR.Name) {
			validationErrors.Set("name", fmt.Errorf("Name: %v already used", sR.Name))
			break
		}
	}

	return validationErrors
}

func (s SkillService) CreateSkill(request entities.SkillRequest) (*entities.Skill, error) {
	validationErrors := s.ValidateSkillRequest(request)
	if validationErrors != nil {
		return nil, errors.Join(BadRequest, validationErrors)
	}

	sR, err := s.repo.Create(*entities.NewSkill(request.Name, request.Description, entities.SkillType(request.Type), request.Importance))

	// TODO: Test if we can create the team directly with skill and meber do i need to add it afterwards
	return sR, err
}

func (s SkillService) GetSkills() (*[]entities.Skill, error) {
	// IDEA: Filter based on user Role? RBAC
	return s.repo.GetAll()
}

func (s SkillService) GetSkillByID(id string) (*entities.Skill, error) {
	var validationErrors errx.ErrorMap
	fid, err := strconv.Atoi(id)
	if err != nil {
		validationErrors.Set("id", err)
	}

	if validationErrors != nil {
		return nil, errors.Join(BadRequest, validationErrors)
	}

	m, err := s.repo.GetByID(uint(fid))
	if err != nil {
		return nil, errors.Join(InternalError, err)
	}
	return m, nil
}

func (s SkillService) DeleteSkill(id string) error {
	var validationErrors errx.ErrorMap
	fid, err := strconv.Atoi(id)
	if err != nil {
		validationErrors.Set("id", err)
	}

	if validationErrors != nil {
		return errors.Join(BadRequest, validationErrors)
	}

	sk, err := s.repo.GetByID(uint(fid))
	if err != nil {
		return errors.Join(InternalError, err)
	}

	if sk == nil {
		return errors.Join(BadRequest, fmt.Errorf("Skill with id: %b, dose't exists", fid))
	}

	err = s.repo.Delete(uint(fid))
	if err != nil {
		return errors.Join(InternalError, err)
	}

	return nil
}

func (s SkillService) UpdateSkill(id string, req entities.SkillRequest) (*entities.Skill, error) {
	// we don't want if name is already used - we know this
	validationErrors := req.ValidateFields()

	fid, err := strconv.Atoi(id)
	if err != nil {
		validationErrors.Set("id", err)
	}

	if validationErrors != nil {
		return nil, errors.Join(BadRequest, validationErrors)
	}

	oldT, err := s.repo.GetByID(uint(fid))
	if err != nil {
		return nil, errors.Join(InternalError, err)
	}

	nM, change := oldT.HasChanges(req.Name, req.Description, entities.SkillType(req.Type), req.Importance)
	if change {
		_, err = s.repo.Update(*nM)
		if err != nil {
			return nil, errors.Join(InternalError, err)
		}
		return nM, nil

	}
	return nil, errors.Join(BadRequest, fmt.Errorf("No changes to Skill: %v found.", id))
}
