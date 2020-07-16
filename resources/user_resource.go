package resources

import (
	"context"

	"github.com/bagus2x/new-sirius/domain"
	"github.com/bagus2x/new-sirius/repositories"
	"github.com/gin-gonic/gin"

	"github.com/bagus2x/new-sirius/services"
	"go.mongodb.org/mongo-driver/mongo"
)

// UserResource -
type UserResource struct {
	userService domain.UserService
}

// NewUserResource -
func NewUserResource(ctx context.Context, db *mongo.Database, r *gin.RouterGroup) {
	ur := repositories.NewUserRepository(ctx, db)
	us := services.NewUserService(ur)
	urs := UserResource{us}
	{
		r.POST("/users/signup", urs.Signup)
		r.POST("/users/signin", urs.Signin)
	}
}

// Signup -
func (urs UserResource) Signup(c *gin.Context) {

}

// Signin -
func (urs UserResource) Signin(c *gin.Context) {

}
