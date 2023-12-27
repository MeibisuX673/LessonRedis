package userController

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/MeibisuX673/LessonRedis/app/controller/dto"
	"github.com/MeibisuX673/LessonRedis/app/model/user"
	"github.com/MeibisuX673/LessonRedis/app/service/securityService"
	"github.com/MeibisuX673/LessonRedis/config/database"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type UserController struct{}

func (uc *UserController) GetItem(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "id должно быть числом",
		})
		return
	}

	cacheUserCmd := database.AllDatabases.DbCache.Db.Get(context.Background(), c.Request.RequestURI)

	if cacheUserCmd.Val() != "" {
		var cacheUser user.User

		json.Unmarshal([]byte(cacheUserCmd.Val()), &cacheUser)

		c.JSON(http.StatusOK, gin.H{
			"user": cacheUser,
		})
		return
	}

	userCmd := database.AllDatabases.DbUser.Db.Get(context.Background(), strconv.Itoa(id))

	if userCmd.Val() == "" {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Пользователь не найден",
		})
		return
	}

	database.AllDatabases.DbCache.Db.Set(context.Background(), c.Request.RequestURI, userCmd.Val(), 0)

	var userInBd user.User

	json.Unmarshal([]byte(userCmd.Val()), &userInBd)

	c.JSON(http.StatusOK, gin.H{
		"user": userInBd,
	})
}

func (uc *UserController) Update(c *gin.Context) {

	currentUser := securityService.GetCurrentUser(c)

	id := c.Param("id")

	_, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "id должно быть числом",
		})
		return
	}

	if currentUser.Id != id {
		c.JSON(http.StatusForbidden, gin.H{
			"message": "Нельзя редактировать чужой профиль",
		})
		return
	}

	userCmd := database.AllDatabases.DbUser.Db.Get(context.Background(), id)

	if userCmd.Val() == "" {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Пользователь не найден",
		})
		return
	}

	var updateUserDto dto.UpdateUser

	c.BindJSON(&updateUserDto)

	currentUser.Email = updateUserDto.Email

	fmt.Println(c.Request.RequestURI)

	database.AllDatabases.DbUser.Db.Set(context.Background(), id, &currentUser, 0)
	database.AllDatabases.DbCache.Db.Set(context.Background(), c.Request.RequestURI, &currentUser, 0)

}
