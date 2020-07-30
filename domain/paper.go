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
	Results     []Result           `json:"results" bson:"results"`
}

// Question - sub paper
type Question struct {
	QstID    int      `json:"qstID" bson:"qstID"`
	Question string   `json:"question" bson:"question"`
	Category string   `json:"category" bson:"category"`
	Key      string   `json:"key" bson:"key"`
	Options  []Option `json:"options" bson:"options"`
}

// Option - sub question
type Option struct {
	OptID  string `json:"optID" bson:"optID"`
	Option string `json:"option" bson:"option"`
	Image  string `json:"image" bson:"image"`
}

// Result - sub paper
type Result struct {
	ID        primitive.ObjectID `json:"_id" bson:"_id"`
	StudentID primitive.ObjectID `json:"student_id" bson:"student_id"`
	Selected  []Selected         `json:"selected" bson:"selected"`
}

// Selected - sub result
type Selected struct {
	QstID  int    `json:"qstID" bson:"qstID"`
	Option string `josn:"option" bson:"option"`
}

// PaperRepository -
type PaperRepository interface {
	InsertOne(p *Paper) (err error)
	FindByID(id string) (res *Paper, err error)
	GetPaper(id string) (res []map[string]interface{}, err error)
	PushExamResult(id string, rst *Result) (resid string, err error)
	GetExamResult(id string, resid string) (res []map[string]interface{}, err error)
}

// PaperService -
type PaperService interface {
	Create(p *Paper) (err error)
	FindByID(id string) (res *Paper, err error)
	GetOneByID(id string) (res map[string]interface{}, err error)
	PushExamResult(id string, rst *Result) (resid string, err error)
	GetExamResult(id string, resid string) (res map[string]interface{}, err error)
}
