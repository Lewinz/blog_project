package blog

import (
	"blog_project/components/db"
	"blog_project/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Create Blog
func Create(c *gin.Context) {
	var (
		blog models.Blog
		db   = db.Get()
	)

	c.ShouldBindJSON(&blog)

	fmt.Println("blog---", blog)

	result := db.Create(&blog)

	if result.RowsAffected > 0 {
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("create blog succeed:%v", blog.ID),
		})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{
			"errMassage": fmt.Sprintf("create blog Faild:%s", result.Error),
		})
	}
}
