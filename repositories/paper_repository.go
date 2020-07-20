package repositories

import (
	"context"

	"github.com/bagus2x/sirius-c/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// PaperRepository -
type PaperRepository struct {
	ctx context.Context
	db  *mongo.Database
}

// NewPaperRepository -
func NewPaperRepository(ctx context.Context, db *mongo.Database) domain.PaperRepository {
	return &PaperRepository{ctx, db}
}

// InsertOne -
func (pr PaperRepository) InsertOne(p *domain.Paper) (err error) {
	p.ID = primitive.NewObjectID()
	for i, v := range p.Questions {
		v.QstID = i
	}
	_, err = pr.db.Collection("paper").InsertOne(pr.ctx, p)
	return err
}

// FindByID -
func (pr PaperRepository) FindByID(id string) (res *domain.Paper, err error) {
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return &domain.Paper{}, err
	}
	err = pr.db.Collection("paper").FindOne(pr.ctx, bson.M{"_id": _id}).Decode(&res)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return &domain.Paper{}, domain.ErrPaperIDNotFound
		}
		return &domain.Paper{}, err
	}
	return res, err
}
