package service

import (
	"errors"

	"github.com/pulsone21/powner/internal/errx"
)

var (
	BadRequest    = errors.New("BadRequest")
	InternalError = errors.New("InternalError")
)

type ServiceErrors struct {
	validationErrors errx.ErrorMap
	err              error
}

func (s *ServiceErrors) Error() string {
	if s.validationErrors != nil {
		return errors.Join(BadRequest, s.validationErrors).Error()
	}
	return errors.Join(InternalError, s.err).Error()
}

func (s *ServiceErrors) GetValidationErrors() *errx.ErrorMap {
	return &s.validationErrors
}
