package service

import "github.com/pkg/errors"

var (
	BadRequest    = errors.New("BadRequest")
	InternalError = errors.New("InternalError")
)
