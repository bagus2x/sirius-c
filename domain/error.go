package domain

import (
	"errors"
)

var (
	// ErrUserAlreadyExist -
	ErrUserAlreadyExist = errors.New("User Already Exist")
	//ErrEmailNotFound -
	ErrEmailNotFound = errors.New("Email not found")
	// ErrInvalidPassword -
	ErrInvalidPassword = errors.New("Invalid password")
)
