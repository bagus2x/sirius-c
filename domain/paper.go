package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Paper -
type Paper struct {
	ID          primitive.ObjectID `json:"_id" bson:"_id"`
	TeacherID   primitive.ObjectID `json:"teacher_id" bson:"teacher_id"`
	Title       string             `json:"title" bson:"title"`
	Description string             `json:"description" bson:"description"`
	Subject     string             `json:"subject" bson:"subject"`
	StartFrom   int64              `json:"startFrom" bson:"startFrom"`
	EndAt       int64              `json:"endAt" bson:"endAt"`
	Questions   []*Question        `json:"questions" bson:"questions"`
}

// Question -
type Question struct {
	QstID    int      `json:"qst_id" bson:"qst_id"`
	Question string   `json:"question" bson:"question"`
	Category string   `json:"category" bson:"category"`
	Key      int      `json:"key" bson:"key"`
	Options  []Option `json:"options" bson:"options"`
}

// Option -
type Option struct {
	OptID  int    `json:"opt_id" bson:"opt_id"`
	Option string `json:"option" bson:"option"`
}

// PaperRepository -
type PaperRepository interface {
	InsertOne(p *Paper) (err error)
	FindByID(id string) (res *Paper, err error)
}

// PaperService -
type PaperService interface {
	Create(p *Paper) (err error)
	FindByID(id string) (res *Paper, err error)
}
