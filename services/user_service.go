package services

import (
	"github.com/bagus2x/new-sirius/domain"
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
func (us UserService) Signup(user domain.User) {
	us.userRepository.GetCountByEmail(user.Email)
	us.userRepository.Create(user)
}

// Signin -
func (us UserService) Signin(email string, password string) {

}
