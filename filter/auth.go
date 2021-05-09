package filter

import (
	"blog_project/models"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// MyClaims is token
type MyClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// TokenExpireDuration token exprire time
const TokenExpireDuration = time.Hour * 24

// MySecret token Secret
var MySecret = []byte("zxl")

// AuthCheck check
func AuthCheck(c *gin.Context) {
	tokenStr, err := c.Cookie("blog_project")

	if err != nil || len(tokenStr) == 0 {
		tokenStr = strings.TrimPrefix(c.GetHeader("Authorization"), "Bearer ")
	}

	if len(tokenStr) == 0 {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	token, err := jwt.ParseWithClaims(tokenStr, &MyClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return MySecret, nil
	})

	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	if _, ok := token.Claims.(*MyClaims); ok && token.Valid { // 校验token
		c.Next()
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
}

// CreateToken create jwt token by name password
func CreateToken(c *gin.Context) {
	var u models.User

	c.ShouldBind(&u)

	if u.Name != "" && u.Password != "" {
		claims := MyClaims{
			u.Name,
			jwt.StandardClaims{
				ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(),
				Issuer:    "blog_project",
			},
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

		result, _ := token.SignedString(MySecret)

		log.Println("token-", result)

		c.SetCookie("blog_project", result, 20000, "/", "", false, true)
		c.JSON(http.StatusOK, gin.H{
			"token": result,
		})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{
			"errMessage": "create token faild: name & password must not empty",
		})
	}
}
