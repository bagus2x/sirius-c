package test

import (
	"context"
	"log"
	"testing"
	"time"

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
			{Question: "Siapakah Hokage pertama?", Key: 1, Options: []domain.Option{{0, "Tobirama"}, {1, "Sandaime"}}},
			{Question: "Siapakah Hokage Kedua?", Key: 1, Options: []domain.Option{{0, "Hiruze"}, {1, "Kabuto"}}},
			{Question: "Siapakah Hokage ketiga?", Key: 1, Options: []domain.Option{{0, "Sasuke"}, {1, "Sandaime"}}},
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
