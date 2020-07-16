package repositories

import (
	"context"
	"log"
	"time"

	"github.com/bagus2x/new-sirius/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

// UserRepository -
type UserRepository struct {
	ctx context.Context
	db  *mongo.Database
}

// NewUserRepository -
func NewUserRepository(ctx context.Context, db *mongo.Database) domain.UserRepository {
	return &UserRepository{ctx, db}
}

// Create -
func (ur UserRepository) Create(user domain.User) (err error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	user.Password = string(hash)
	user.ID = primitive.NewObjectID()
	user.CreatedAt = time.Now().Unix()
	if err != nil {
		log.Fatal(err)
		return err
	}
	_, err = ur.db.Collection("user").InsertOne(ur.ctx, user)
	return err
}

// FindByEmailAndPassword -
func (ur UserRepository) FindByEmailAndPassword(email string, password string) (user domain.User, err error) {
	err = ur.db.Collection("user").FindOne(ur.ctx, bson.M{"email": email}).Decode(&user)
	if err != nil {
		return user, domain.ErrEmailNotFound
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return user, domain.ErrInvalidPassword
	}
	return user, nil
}

// GetCountByEmail -
func (ur UserRepository) GetCountByEmail(email string) (c int64, err error) {
	return ur.db.Collection("user").CountDocuments(ur.ctx, bson.M{"email": email})
}
