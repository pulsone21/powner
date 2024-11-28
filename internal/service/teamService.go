package service

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/pulsone21/powner/internal/entities"
	"github.com/pulsone21/powner/internal/errx"
	"github.com/pulsone21/powner/internal/repos"
)

type TeamService struct {
	repo repos.TeamRepository
}

func (s TeamService) ValidateTeamRequest(t entities.TeamRequest) errx.ErrorMap {
	validationErrors := t.ValidateFields()
	ts, err := s.repo.GetAll()
	if err != nil {
		validationErrors.Set("Team", errors.Join(fmt.Errorf("could not fetch teams from repository"), err))
	}

	for _, te := range *ts {
		if te.Name == t.Name {
			validationErrors.Set("name", fmt.Errorf("name is already used."))
		}
	}

	return validationErrors
}

func (s TeamService) CreateTeam(request entities.TeamRequest) (*entities.Team, error) {
	validationErrors := s.ValidateTeamRequest(request)
	if validationErrors != nil {
		return nil, errors.Join(BadRequest, validationErrors)
	}

	t, err := s.repo.Create(*entities.NewTeam(request.Name, request.Description))
	// TODO: Test if we can create the team directly with skill and meber do i need to add it afterwards
	return t, err
}

func (s TeamService) GetTeams() (*[]entities.Team, error) {
	// IDEA: Filter based on user Role? RBAC
	return s.repo.GetAll()
}

func (s TeamService) GetTeamByID(id string) (*entities.Team, error) {
	var validationErrors errx.ErrorMap
	fid, err := strconv.Atoi(id)
	if err != nil {
		validationErrors.Set("id", err)
	}

	if validationErrors != nil {
		return nil, errors.Join(BadRequest, validationErrors)
	}
	t, err := s.repo.GetByID(uint(fid))
	if err != nil {
		return nil, errors.Join(InternalError, err)
	}
	return t, nil
}

func (s TeamService) DeleteTeam(id string) error {
	var validationErrors errx.ErrorMap
	fid, err := strconv.Atoi(id)
	if err != nil {
		validationErrors.Set("id", err)
	}

	if validationErrors != nil {
		return errors.Join(BadRequest, validationErrors)
	}

	t, err := s.repo.GetByID(uint(fid))
	if err != nil {
		return errors.Join(InternalError, err)
	}

	if t == nil {
		return errors.Join(BadRequest, fmt.Errorf("Team with id: %b dosen't exisits", fid))
	}

	err = s.repo.Delete(*t)
	if err != nil {
		return errors.Join(InternalError, err)
	}

	return nil
}

func (s TeamService) UpdateTeam(id string, request entities.TeamRequest) (*entities.Team, error) {
	validationErrors := request.ValidateFields()

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
	if oldT == nil {
		return nil, errors.Join(BadRequest, fmt.Errorf("Team with id: %b doesn't exist", fid))
	}

	nT, change := oldT.HasChanges(request.Name, request.Description, &oldT.Skills, &oldT.Members)
	if change {
		newTeam, err := s.repo.Update(*nT)
		if err != nil {
			return nil, errors.Join(InternalError, err)
		}
		return newTeam, nil
	}
	return nil, errors.Join(BadRequest, fmt.Errorf("No changes to Team: %v found.", id))
}
