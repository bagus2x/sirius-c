package test

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/bagus2x/sirius-c/db"
	"github.com/bagus2x/sirius-c/domain"
	"github.com/bagus2x/sirius-c/repositories"
)

func TestCreatePaper(t *testing.T) {
	db, cancel, err := db.Connect(dbURI, dbName, 10*time.Second)
	defer cancel()
	if err != nil {
		log.Fatal(err)
	}
	pr := repositories.NewPaperRepository(context.TODO(), db)
	err = pr.InsertOne(&domain.Paper{
		Title:       "Ujian Chunin",
		Description: "Ini adalah ujian chunin",
		StartFrom:   time.Now().Unix(),
		EndAt:       time.Now().Add(time.Hour * 2).Unix(),
		Questions: []*domain.Question{
			{Question: "Siapakah Hokage pertama?", Key: "A", Options: []domain.Option{{"A", "Tobirama", domain.Image{"", "200px"}}, {"B", "Sandaime", domain.Image{"", "200px"}}}},
			{Question: "Siapakah Hokage Keduaa?", Key: "B", Options: []domain.Option{{"A", "Hiruze", domain.Image{"", "200px"}}, {"B", "Kabuto", domain.Image{"", "200px"}}}},
			{Question: "Siapakah Hokage ketiga?", Key: "C", Options: []domain.Option{{"A", "Sasuke", domain.Image{"", "200px"}}, {"B", "Sandaime", domain.Image{"", "200px"}}}},
		},
	},
	)
	if err != nil {
		log.Fatal(err)
	}
}
func TestFindPaperByID(t *testing.T) {
	db, cancel, err := db.Connect(dbURI, dbName, 10*time.Second)
	defer cancel()
	if err != nil {
		log.Fatal(err)
	}
	pr := repositories.NewPaperRepository(context.TODO(), db)
	res, err := pr.FindByID("5f13d3da33800aff4d1cf149")
	if err != nil {
		log.Fatal(err)
	}
	t.Log(res)
}

func TestGetOneBydID(t *testing.T) {
	db, cancel, err := db.Connect(dbURI, dbName, 10*time.Second)
	defer cancel()
	if err != nil {
		log.Fatal(err)
	}
	pr := repositories.NewPaperRepository(context.TODO(), db)
	res, err := pr.GetPaper("5f1e7312d0802c4eb0c5548c")
	if err != nil {
		log.Fatal(err)
		return
	}
	res2, err := json.Marshal(res)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(string(res2))
}

func TestPushExamResult(*testing.T) {
	db, cancel, err := db.Connect(dbURI, dbName, 10*time.Second)
	defer cancel()
	if err != nil {
		log.Fatal(err)
	}
	stdID, _ := primitive.ObjectIDFromHex("5f1704e65191127f76aced7e")
	pr := repositories.NewPaperRepository(context.TODO(), db)
	resid, err := pr.PushExamResult("5f1fcf54efcbf0ec9ba24b49", &domain.Result{StudentID: stdID, Selected: []domain.Selected{{QstID: 0, Option: ""}, {QstID: 1, Option: ""}, {QstID: 2, Option: ""}}})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resid)
}

func TestGetExamResult(t *testing.T) {
	db, cancel, err := db.Connect(dbURI, dbName, 10*time.Second)
	defer cancel()
	if err != nil {
		log.Fatal(err)
	}
	pr := repositories.NewPaperRepository(context.TODO(), db)
	res, err := pr.GetExamResult("5f22e484d5c048695e05b01d", "5f22e6fcd5c048695e05b020")
	res2, _ := json.Marshal(res)
	fmt.Println(string(res2))
	if err != nil {
		log.Fatal(err)
	}
}

func TestEncode(t *testing.T) {
	// "studentid" : ObjectId("5f1704e65191127f76aced7e"),
	raw := `{
		"student_id" : "5f1704e65191127f76aced7a",
		"selected" : [
			{
				"qstID" : 1,
				"option" : "A"
			}
		]
	}`
	var to domain.Result
	err := json.Unmarshal([]byte(raw), &to)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println(to)
}
