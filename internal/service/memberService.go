package service

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/pulsone21/powner/internal/entities"
	"github.com/pulsone21/powner/internal/errx"
)

func (s Service) ValidateMemberRequest(t entities.MemberRequest) *errx.ErrorMap {
	validationErrors := t.ValidateFields()

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

	return &validationErrors
}

func (s Service) CreateMember(request entities.MemberRequest) (*entities.Member, error) {
	validationErrors := s.ValidateMemberRequest(request)
	if validationErrors != nil {
		return nil, errors.Join(BadRequest, validationErrors)
	}

	t, err := s.MemberRepo.Create(*entities.NewMember(request.Name, request.Age))
	// TODO: Test if we can create the team directly with skill and meber do i need to add it afterwards
	return t, err
}

func (s Service) GetMembers() (*[]entities.Member, error) {
	// IDEA: Filter based on user Role? RBAC
	return s.MemberRepo.GetAll()
}

func (s Service) GetMemberByID(id string) (*entities.Member, error) {
	var validationErrors errx.ErrorMap
	fid, err := strconv.Atoi(id)
	if err != nil {
		validationErrors.Set("id", err)
	}

	if validationErrors != nil {
		return nil, errors.Join(BadRequest, validationErrors)
	}

	m, err := s.MemberRepo.GetByID(uint(fid))
	if err != nil {
		return nil, errors.Join(InternalError, err)
	}
	return m, nil
}

func (s Service) DeleteMember(id string) error {
	var validationErrors errx.ErrorMap
	fid, err := strconv.Atoi(id)
	if err != nil {
		validationErrors.Set("id", err)
	}

	if validationErrors != nil {
		return errors.Join(BadRequest, validationErrors)
	}
	err = s.MemberRepo.Delete(uint(fid))
	if err != nil {
		return errors.Join(InternalError, err)
	}

	return nil
}

func (s Service) UpdateMember(id string, request entities.MemberRequest) (*entities.Member, error) {
	validationErrors := s.ValidateMemberRequest(request)

	fid, err := strconv.Atoi(id)
	if err != nil {
		validationErrors.Set("id", err)
	}

	if validationErrors != nil {
		return nil, errors.Join(BadRequest, validationErrors)
	}

	oldT, err := s.MemberRepo.GetByID(uint(fid))
	if err != nil {
		return nil, errors.Join(InternalError, err)
	}

	nM, change := oldT.HasChanges(request.Name, request.Age)
	if change {
		_, err = s.MemberRepo.Update(*nM)
		return nM, errors.Join(InternalError, err)
	}
	return nil, errors.Join(BadRequest, fmt.Errorf("No changes to Member: %v found.", id))
}
