package domain

import (
	"errors"
)

var (
	// ErrEmailAlreadyExist -
	ErrEmailAlreadyExist = errors.New("Email already exist")
	//ErrEmailNotFound -
	ErrEmailNotFound = errors.New("Email not found")
	// ErrInvalidPassword -
	ErrInvalidPassword = errors.New("Invalid password")
)
