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

func (s SkillService) CreateSkill(request entities.SkillRequest) (*entities.Skill, *ServiceErrors) {
	validationErrors := s.ValidateSkillRequest(request)
	if validationErrors != nil {
		return nil, &ServiceErrors{validationErrors: validationErrors}
	}

	sR, err := s.repo.Create(*entities.NewSkill(request.Name, request.Description, entities.SkillType(request.Type), request.Importance))
	// TODO: Test if we can create the team directly with skill and meber do i need to add it afterwards
	if err != nil {
		return nil, &ServiceErrors{err: err}
	}

	return sR, nil
}

func (s SkillService) GetSkills() (*[]entities.Skill, *ServiceErrors) {
	// IDEA: Filter based on user Role? RBAC
	sk, err := s.repo.GetAll()
	if err != nil {
		return nil, &ServiceErrors{err: err}
	}
	return sk, nil
}

func (s SkillService) GetSkillByID(id string) (*entities.Skill, *ServiceErrors) {
	var validationErrors errx.ErrorMap
	fid, err := strconv.Atoi(id)
	if err != nil {
		validationErrors.Set("id", err)
	}

	if validationErrors != nil {
		return nil, &ServiceErrors{validationErrors: validationErrors}
	}

	m, err := s.repo.GetByID(uint(fid))
	if err != nil {
		return nil, &ServiceErrors{err: errors.Join(InternalError, err)}
	}
	return m, nil
}

func (s SkillService) DeleteSkill(id string) *ServiceErrors {
	var validationErrors errx.ErrorMap
	fid, err := strconv.Atoi(id)
	if err != nil {
		validationErrors.Set("id", err)
	}

	if validationErrors != nil {
		return &ServiceErrors{validationErrors: validationErrors}
	}

	sk, err := s.repo.GetByID(uint(fid))
	if err != nil {
		return &ServiceErrors{err: errors.Join(InternalError, err)}
	}

	if sk == nil {
		return &ServiceErrors{err: errors.Join(BadRequest, fmt.Errorf("Skill with id: %b, dose't exists", fid))}
	}

	err = s.repo.Delete(uint(fid))
	if err != nil {
		return &ServiceErrors{err: errors.Join(InternalError, err)}
	}

	return nil
}

func (s SkillService) UpdateSkill(id string, req entities.SkillRequest) (*entities.Skill, *ServiceErrors) {
	// we don't want if name is already used - we know this
	validationErrors := req.ValidateFields()

	fid, err := strconv.Atoi(id)
	if err != nil {
		validationErrors.Set("id", err)
	}

	if validationErrors != nil {
		return nil, &ServiceErrors{validationErrors: validationErrors}
	}

	oldT, err := s.repo.GetByID(uint(fid))
	if err != nil {
		return nil, &ServiceErrors{err: errors.Join(InternalError, err)}
	}

	nM, change := oldT.HasChanges(req.Name, req.Description, entities.SkillType(req.Type), req.Importance)
	if change {
		_, err = s.repo.Update(*nM)
		if err != nil {
			return nil, &ServiceErrors{err: errors.Join(InternalError, err)}
		}
		return nM, nil

	}
	return nil, &ServiceErrors{err: errors.Join(BadRequest, fmt.Errorf("No changes to Skill: %v found.", id))}
}
