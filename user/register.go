package user

import (
	"blog_project/components/db"
	"blog_project/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// createArg create user arg
type CreateArg struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email" gorm:"size:128;default:null"`
	Phone    string `json:"phone"`
}

// Create register blog
func Create(c *gin.Context) {
	var (
		arg = CreateArg{}
		db  = db.Get()
	)

	c.ShouldBindJSON(&arg)

	var user = models.User{
		Name:     arg.Name,
		Password: arg.Password,
		Email:    arg.Email,
		Phone:    arg.Phone,
		Role:     models.RoleConsumer,
	}

	result := db.Create(&user)
	if result.RowsAffected > 0 {
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("user create succeed,userName:%s", user.Name),
		})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{
			"errMessage": result.Error,
		})
	}
}
