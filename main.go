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
	cancel, err := db.NewConnection.Init(dbURI, dbName, 10*time.Second)
	defer cancel()
	if err != nil {
		log.Fatal(err)
	}
	r.GET("/", func(c *gin.Context) {
		c.String(200, "hello, world")
	})
	// Setup Routes
	rg := r.Group("/api")
	// [GET, POST] api/users/
	resources.NewUserResource(rg)
	// [GET, POST] api/papers/
	resources.NewPaperResource(rg)
	r.Run(":" + port)
}
