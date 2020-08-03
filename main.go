package main

import (
	"log"
	"os"
	"time"

	"github.com/bagus2x/sirius-c/middleware"

	"github.com/bagus2x/sirius-c/db"
	"github.com/bagus2x/sirius-c/resources"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

var (
	port   = os.Getenv("PORT")
	dbURI  = os.Getenv("DB_URI")
	dbName = os.Getenv("DB_NAME")
)

func main() {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(gzip.Gzip(gzip.DefaultCompression))
	r.Use(middleware.CORSMiddleware())
	// Setup Database
	db, cancel, err := db.Connect(dbURI, dbName, 10*time.Second)
	defer cancel()
	if err != nil {
		log.Fatal(err)
	}
	// Setup Routes
	rg := r.Group("/api")
	// [GET, POST] api/users/
	resources.NewUserResource(db, rg)
	// [GET, POST] api/papers/
	resources.NewPaperResource(db, rg)
	r.Run(":" + port)
}
