package securityService

import (
	"context"
	"encoding/json"
	"github.com/MeibisuX673/LessonRedis/app/model/user"
	"github.com/MeibisuX673/LessonRedis/config/database"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"os"
	"strings"
)

func GetCurrentUser(c *gin.Context) user.User {

	tokenString := c.GetHeader("Authorization")

	tokenData := strings.Fields(tokenString)
	token, _ := jwt.Parse(tokenData[1], func(token *jwt.Token) (interface{}, error) {

		return []byte(os.Getenv("SECRET")), nil
	})

	var currentUser user.User

	id := token.Claims.(jwt.MapClaims)["sub"].(string)

	userCmd := database.AllDatabases.DbUser.Db.Get(context.Background(), id)

	json.Unmarshal([]byte(userCmd.Val()), &currentUser)

	return currentUser

}
