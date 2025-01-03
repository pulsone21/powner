package service

import (
	"errors"
	"fmt"
	"log"
	"strconv"

	"github.com/pulsone21/powner/internal/entities"
	"github.com/pulsone21/powner/internal/errx"
	"github.com/pulsone21/powner/internal/repos"
)

type MemberService struct {
	repo repos.MemberRepository
}

func (s MemberService) CreateMember(request entities.MemberRequest) (*entities.Member, *ServiceErrors) {
	validationErrors := request.ValidateFields()
	log.Printf("Member Req Validation")
	if validationErrors.Error() != "<nil>" {
		return nil, &ServiceErrors{validationErrors: validationErrors}
	}
	log.Printf("Member Req Validation passed")

	m, err := s.repo.Create(*entities.NewMember(request.Name, request.Age))
	if err != nil {
		return nil, &ServiceErrors{err: err}
	}

	return m, nil
}

func (s MemberService) GetMembers() (*[]entities.Member, *ServiceErrors) {
	mems, err := s.repo.GetAll()
	if err != nil {
		return nil, &ServiceErrors{err: err}
	}

	// IDEA: Filter based on user Role? RBAC

	return mems, nil
}

func (s MemberService) GetMemberByID(id string) (*entities.Member, *ServiceErrors) {
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
		return nil, &ServiceErrors{err: err}
	}
	return m, nil
}

func (s MemberService) DeleteMember(id string) *ServiceErrors {
	var validationErrors errx.ErrorMap
	fid, err := strconv.Atoi(id)
	if err != nil {
		validationErrors.Set("id", err)
	}

	if validationErrors != nil {
		return &ServiceErrors{validationErrors: validationErrors}
	}
	err = s.repo.Delete(uint(fid))
	if err != nil {
		return &ServiceErrors{err: err}
	}

	return nil
}

func (s MemberService) UpdateMember(id string, request entities.MemberRequest) (*entities.Member, *ServiceErrors) {
	validationErrors := request.ValidateFields()

	fid, err := strconv.Atoi(id)
	if err != nil {
		validationErrors.Set("id", err)
	}

	if validationErrors != nil {
		return nil, &ServiceErrors{validationErrors: validationErrors}
	}

	oldT, err := s.repo.GetByID(uint(fid))
	if err != nil {
		return nil, &ServiceErrors{err: err}
	}

	if oldT == nil {
		return nil, &ServiceErrors{err: errors.Join(BadRequest, fmt.Errorf("Member with id: %v not found", fid))}
	}

	nM, change := oldT.HasChanges(request.Name, request.Age)
	if change {
		m, err := s.repo.Update(*nM)
		if err != nil {
			return nil, &ServiceErrors{err: err}
		}
		return m, nil
	}

	return nil, &ServiceErrors{err: errors.Join(BadRequest, fmt.Errorf("No changes to Member: %v found.", id))}
}
