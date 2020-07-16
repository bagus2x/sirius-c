package domain

import (
	"errors"
)

var (
	// ErrEmailAlreadyExist -
	ErrEmailAlreadyExist = errors.New("Email Already Exist")
	//ErrEmailNotFound -
	ErrEmailNotFound = errors.New("Email not found")
	// ErrInvalidPassword -
	ErrInvalidPassword = errors.New("Invalid password")
)
