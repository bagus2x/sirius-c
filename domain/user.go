package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User -
type User struct {
	ID        primitive.ObjectID `json:"_id" bson:"_id"`
	Username  string             `json:"Username" bson:"Username" validate:"required"`
	Fullname  string             `json:"fullName" bson:"fullName" validate:"required"`
	Email     string             `json:"email" bson:"email" validate:"required,email"`
	Password  string             `json:"password" bson:"password" validate:"required"`
	Role      string             `json:"role" bson:"role" validate:"required"`
	CreatedAt int64              `json:"createdAt" bson:"createdAt"`
	UpdatedAt int64              `json:"updatedAt" bson:"updatedAt"`
}

// UserRepository -
type UserRepository interface {
	FindByEmailAndPassword(email string, password string) (user User, err error)
	Create(user User) (err error)
	GetCountByEmail(email string) (c int64, err error)
	// FindById(id string) User
}

// UserService -
type UserService interface {
	Signup(user User)
	Signin(email string, password string)
}
