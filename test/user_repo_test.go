package test

import (
	"context"
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/bagus2x/new-sirius/domain"

	"github.com/bagus2x/new-sirius/db"
	"github.com/bagus2x/new-sirius/repositories"
)

func TestCreate(t *testing.T) {
	db, cancel, err := db.Connect(dbURI, dbName, 10*time.Second)
	defer cancel()
	if err != nil {
		log.Fatal(err)
	}
	ur := repositories.NewUserRepository(context.TODO(), db)
	err = ur.Create(domain.User{
		Username: "bagus_kece",
		Password: "bagus",
		Email:    "bagus@gmail.com",
		Fullname: "tubagus",
	})
	if err != nil {
		log.Fatal(err)
	}
}

func TestFindByIdAndEmail(t *testing.T) {
	db, cancel, err := db.Connect(dbURI, dbName, 10*time.Second)
	defer cancel()
	if err != nil {
		log.Fatal(err)
	}
	ur := repositories.NewUserRepository(context.TODO(), db)
	_, err = ur.FindByEmailAndPassword("robi@gmail.com", "robi")
	if err != nil {
		log.Fatal(err)
	}
}

func TestGetCountByEmail(t *testing.T) {
	db, cancel, err := db.Connect(dbURI, dbName, 10*time.Second)
	defer cancel()
	if err != nil {
		log.Fatal(err)
	}
	ur := repositories.NewUserRepository(context.TODO(), db)
	count, _ := ur.GetCountByEmail("arobi@gmail.com")
	t.Log(count)

}
func TestHello(t *testing.T) {
	fmt.Println("hello")
}
