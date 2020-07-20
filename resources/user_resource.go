package resources

import (
	"context"
	"net/http"
	"strings"

	"github.com/bagus2x/sirius-c/domain"
	"github.com/bagus2x/sirius-c/repositories"
	"github.com/bagus2x/sirius-c/services"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

// UserResource -
type UserResource struct {
	userService domain.UserService
}

// NewUserResource -
func NewUserResource(db *mongo.Database, r *gin.RouterGroup) {
	ur := repositories.NewUserRepository(context.TODO(), db)
	us := services.NewUserService(ur)
	urs := UserResource{us}
	{
		r.POST("/users/signup", urs.Signup)
		r.POST("/users/signin", urs.Signin)
	}
}

// Signup -
func (urs UserResource) Signup(c *gin.Context) {
	var u domain.User
	err := c.BindJSON(&u)
	if err != nil {
		c.JSON(getStatusCodeUser(err), gin.H{"success": false, "message": err.Error()})
		return
	}
	tokStr, err := urs.userService.Signup(&u)
	if err != nil {
		c.JSON(getStatusCodeUser(err), gin.H{"success": false, "message": err.Error()})
		return
	}
	c.JSON(200, gin.H{"success": true, "token": tokStr})
}

// Signin -
func (urs UserResource) Signin(c *gin.Context) {
	var sf struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	c.BindJSON(&sf)
	tokStr, err := urs.userService.Signin(sf.Email, sf.Password)
	if err != nil {
		c.JSON(getStatusCodeUser(err), gin.H{"success": false, "message": err.Error()})
		return
	}
	c.JSON(200, gin.H{"success": true, "token": tokStr})
}

func getStatusCodeUser(err error) int {
	switch err {
	case domain.ErrEmailAlreadyExist:
		return http.StatusConflict
	case domain.ErrEmailNotFound:
		return http.StatusNotFound
	case domain.ErrInvalidPassword:
		return http.StatusUnauthorized
	default:
		if strings.Contains(err.Error(), "Key:") {
			return http.StatusBadRequest
		}
		return http.StatusInternalServerError
	}
}
