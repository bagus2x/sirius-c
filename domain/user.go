package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User -
type User struct {
	ID        primitive.ObjectID `json:"_id" bson:"_id"`
	Username  string             `json:"username" bson:"username" validate:"required,min=4,max=50"`
	Fullname  string             `json:"fullName" bson:"fullName" validate:"required,min=4,max=50"`
	Email     string             `json:"email" bson:"email" validate:"required,email"`
	Password  string             `json:"password" bson:"password" validate:"required,min=8,max=20"`
	Role      string             `json:"role" bson:"role" validate:"required"`
	CreatedAt int64              `json:"createdAt" bson:"createdAt"`
	UpdatedAt int64              `json:"updatedAt" bson:"updatedAt"`
	Validated bool               `json:"validated" bson:"validated"`
}

// UserRepository -
type UserRepository interface {
	FindByEmailAndPassword(email string, password string) (res *User, err error)
	Create(u *User) (err error)
	CountByEmail(email string) (c int64, err error)
	// FindById(id string) Users
}

// UserService -
type UserService interface {
	Signup(u *User) (tokStr string, err error)
	Signin(email string, password string) (tokStr string, err error)
}
