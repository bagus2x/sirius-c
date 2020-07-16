package resources

import (
	"context"
	"net/http"
	"strings"

	"github.com/bagus2x/new-sirius/domain"
	"github.com/bagus2x/new-sirius/repositories"
	"github.com/bagus2x/new-sirius/services"
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
	var user domain.User
	c.BindJSON(&user)
	err := urs.userService.Signup(user)
	if err != nil {
		c.JSON(getStatusCode(err), gin.H{"success": false, "message": err.Error()})
		return
	}
	c.JSON(200, gin.H{"success": true})
}

// Signin -
func (urs UserResource) Signin(c *gin.Context) {

}

func getStatusCode(err error) int {
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
