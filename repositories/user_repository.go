package repositories

import (
	"context"
	"log"
	"time"

	"github.com/bagus2x/sirius-c/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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
func (ur UserRepository) Create(u *domain.User) (err error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), 10)
	u.Password = string(hash)
	u.ID = primitive.NewObjectID()
	u.CreatedAt = time.Now().Unix()
	u.Validated = false
	if err != nil {
		log.Fatal(err)
		return err
	}
	_, err = ur.db.Collection("user").InsertOne(ur.ctx, u)
	return err
}

// FindByEmailAndPassword -
func (ur UserRepository) FindByEmailAndPassword(email string, password string) (res *domain.User, err error) {
	err = ur.db.Collection("user").FindOne(ur.ctx, bson.M{"email": email}).Decode(&res)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return &domain.User{}, domain.ErrEmailNotFound
		}
		return &domain.User{}, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(res.Password), []byte(password))
	if err != nil {
		return &domain.User{}, domain.ErrInvalidPassword
	}
	return res, nil
}

// CountByEmail -
func (ur UserRepository) CountByEmail(email string) (tot int64, err error) {
	opts := options.Count().SetLimit(1)
	return ur.db.Collection("user").CountDocuments(ur.ctx, bson.M{"email": email}, opts)
}
