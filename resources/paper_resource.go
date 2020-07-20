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

// PaperResource -
type PaperResource struct {
	paperService domain.PaperService
}

// NewPaperResource -
func NewPaperResource(db *mongo.Database, r *gin.RouterGroup) {
	pr := repositories.NewPaperRepository(context.TODO(), db)
	ps := services.NewPaperService(pr)
	prc := PaperResource{ps}
	{
		r.GET("/papers/:id", prc.GetByID)
		r.POST("/papers", prc.Create)
	}
}

// Create -
func (pr PaperResource) Create(c *gin.Context) {
	var p domain.Paper
	err := c.BindJSON(&p)
	if err != nil {
		c.JSON(getStatusCodePaper(err), gin.H{"success": false, "message": err.Error()})
		return
	}
	err = pr.paperService.Create(&p)
	if err != nil {
		c.JSON(getStatusCodePaper(err), gin.H{"success": false, "message": err.Error()})
		return
	}
	c.JSON(200, gin.H{"success": true})
	return
}

// GetByID -
func (pr PaperResource) GetByID(c *gin.Context) {
	id := c.Param("id")
	res, err := pr.paperService.FindByID(id)
	if err != nil {
		c.JSON(getStatusCodePaper(err), gin.H{"success": false, "message": err.Error()})
		return
	}
	c.JSON(200, gin.H{"success": true, "paper": res})
}

func getStatusCodePaper(err error) int {
	switch err {
	case domain.ErrPaperIDNotFound:
		return http.StatusNotFound
	default:
		if strings.Contains(err.Error(), "Key:") {
			return http.StatusBadRequest
		}
		return http.StatusInternalServerError
	}
}
