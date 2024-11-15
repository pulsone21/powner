package service

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/pulsone21/powner/internal/entities"
	"github.com/pulsone21/powner/internal/errx"
	"github.com/pulsone21/powner/internal/repos"
)

type MemberService struct {
	repo repos.MemberRepository
}

func (s MemberService) validateMemberRequest(t entities.MemberRequest) *errx.ErrorMap {
	validationErrors := t.ValidateFields()
	return &validationErrors
}

func (s MemberService) CreateMember(request entities.MemberRequest) (*entities.Member, error) {
	validationErrors := s.validateMemberRequest(request)
	if validationErrors != nil {
		return nil, errors.Join(BadRequest, validationErrors)
	}

	t, err := s.repo.Create(*entities.NewMember(request.Name, request.Age))
	return t, err
}

func (s MemberService) GetMembers() (*[]entities.Member, error) {
	mems, err := s.repo.GetAll()
	if err != nil {
		return nil, errors.Join(InternalError, err)
	}

	// IDEA: Filter based on user Role? RBAC

	return mems, nil
}

func (s MemberService) GetMemberByID(id string) (*entities.Member, error) {
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

func (s MemberService) DeleteMember(id string) error {
	var validationErrors errx.ErrorMap
	fid, err := strconv.Atoi(id)
	if err != nil {
		validationErrors.Set("id", err)
	}

	if validationErrors != nil {
		return errors.Join(BadRequest, validationErrors)
	}
	err = s.repo.Delete(uint(fid))
	if err != nil {
		return errors.Join(InternalError, err)
	}

	return nil
}

func (s MemberService) UpdateMember(id string, request entities.MemberRequest) (*entities.Member, error) {
	validationErrors := s.validateMemberRequest(request)

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

	nM, change := oldT.HasChanges(request.Name, request.Age)
	if change {
		_, err = s.repo.Update(*nM)
		return nM, errors.Join(InternalError, err)
	}
	return nil, errors.Join(BadRequest, fmt.Errorf("No changes to Member: %v found.", id))
}
