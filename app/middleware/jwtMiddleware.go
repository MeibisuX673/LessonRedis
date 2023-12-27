package middleware

import (
	"context"
	"encoding/json"
	user2 "github.com/MeibisuX673/LessonRedis/app/model/user"
	"github.com/MeibisuX673/LessonRedis/config/database"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strings"
	"time"
)

func JwtMiddleware(c *gin.Context) {

	tokenString := c.GetHeader("Authorization")

	if len(tokenString) == 0 {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "No Authorized",
		})
		return
	}

	tokenData := strings.Fields(tokenString)

	if tokenData[0] != "Bearer" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid token type",
		})
		return
	}

	token, _ := jwt.Parse(tokenData[1], func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "no authorized",
			})
		}

		return []byte(os.Getenv("SECRET")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Время токена истекло",
			})

		}

		var user user2.User

		id := claims["sub"].(string)

		userCmd := database.AllDatabases.DbUser.Db.Get(context.Background(), id)

		if userCmd.Val() == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Not Unauthorized",
			})
		}

		json.Unmarshal([]byte(userCmd.Val()), &user)

	} else {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Not Unauthorized",
		})

	}

	c.Next()

}
