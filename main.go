package main

import (
	"log"
	"os"
	"time"

	"github.com/bagus2x/new-sirius/resources"

	"github.com/bagus2x/new-sirius/db"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var (
	port   = os.Getenv("DEV_PORT")
	dbURI  = os.Getenv("DB_URI")
	dbName = os.Getenv("DB_NAME")
)

func main() {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowHeaders:    []string{"x-auth-token", "Content-Type"},
	}))
	// Setup Database
	db, cancel, err := db.Connect(dbURI, dbName, 10*time.Second)
	defer cancel()
	if err != nil {
		log.Fatal(err)
	}
	// Setup Routes
	rg := r.Group("/api")
	resources.NewUserResource(db, rg)
	r.Run(":" + port)
}
