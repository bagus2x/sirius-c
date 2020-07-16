package services

import (
	"github.com/bagus2x/new-sirius/domain"
	"github.com/go-playground/validator/v10"
)

// UserService -
type UserService struct {
	userRepository domain.UserRepository
}

// NewUserService -
func NewUserService(userRepository domain.UserRepository) domain.UserService {
	return &UserService{userRepository}
}

// Signup -
func (us UserService) Signup(user domain.User) (err error) {
	count, _ := us.userRepository.CountByEmail(user.Email)
	if count > 0 {
		return domain.ErrEmailAlreadyExist
	}
	v := validator.New()
	err = v.Struct(user)
	if err != nil {
		return err
	}
	err = us.userRepository.Create(user)
	return err
}

// Signin -
func (us UserService) Signin(email string, password string) (err error) {
	return err
}
