package services

import (
	"github.com/bagus2x/sirius-c/domain"
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
func (us UserService) Signup(u *domain.User) (tokStr string, err error) {
	count, _ := us.userRepository.CountByEmail(u.Email)
	if count > 0 {
		return "", domain.ErrEmailAlreadyExist
	}
	v := validator.New()
	err = v.Struct(u)
	if err != nil {
		return "", err
	}
	err = us.userRepository.Create(u)
	if err != nil {
		return "", err
	}
	return CreateToken(u.ID.Hex(), u.Validated)
}

// Signin -
func (us UserService) Signin(email string, password string) (tokStr string, err error) {
	v := validator.New()
	err = v.Struct(struct {
		Email    string `validate:"required,min=4,max=50"`
		Password string `validate:"required,min=8,max=20"`
	}{email, password})
	if err != nil {
		return "", err
	}
	res, err := us.userRepository.FindByEmailAndPassword(email, password)
	if err != nil {
		return "", err
	}
	return CreateToken(res.ID.Hex(), res.Validated)
}
