package authController

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/MeibisuX673/LessonRedis/app/model/user"
	"github.com/MeibisuX673/LessonRedis/app/service/authService/jwtService"
	"github.com/MeibisuX673/LessonRedis/app/service/securityService"
	"github.com/MeibisuX673/LessonRedis/app/util"
	"github.com/MeibisuX673/LessonRedis/config/database"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"time"
)

type AuthController struct {
}

func (ac *AuthController) SignUp(c *gin.Context) {

	var createUser user.User

	if err := c.BindJSON(&createUser); err != nil {
		log.Fatal(err)
	}

	userCmd := database.AllDatabases.DbUser.Db.Get(context.Background(), createUser.Id)

	if userCmd.Val() != "" {
		c.JSON(http.StatusConflict, gin.H{
			"message": "Данный пользователь уже существует",
		})
		return
	}

	code := strconv.Itoa(util.GetRandomNumber())

	if err := database.AllDatabases.DbCode.Db.Set(context.Background(), code, createUser.Id, time.Minute*2).Err(); err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}

	createUser.IsActivate = false

	if err := database.AllDatabases.DbUser.Db.Set(context.Background(), createUser.Id, &createUser, 0).Err(); err != nil {
		log.Fatal(err)
	}

	token, err := jwtService.CreateJwtToken(createUser)
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusCreated, gin.H{
		"code":        code,
		"accessToken": token,
	})
}

func (ac *AuthController) ActivateUser(c *gin.Context) {

	currentUser := securityService.GetCurrentUser(c)

	if currentUser.IsActivate {
		c.JSON(http.StatusOK, gin.H{
			"message": "Аккаунт уже активирован",
		})
		return
	}

	codeStruct := struct {
		Code string
	}{}
	c.BindJSON(&codeStruct)

	codeCmd := database.AllDatabases.DbCode.Db.Get(context.Background(), codeStruct.Code)

	if codeCmd.Val() == "" {
		c.JSON(http.StatusConflict, gin.H{
			"message": "Неверный код. Запросите новый",
		})
		return
	}

	userCmd := database.AllDatabases.DbUser.Db.Get(context.Background(), codeCmd.Val())

	var previousUser user.User

	json.Unmarshal([]byte(userCmd.Val()), &previousUser)

	previousUser.IsActivate = true

	database.AllDatabases.DbUser.Db.Set(context.Background(), codeCmd.Val(), &previousUser, 0)

	c.JSON(http.StatusOK, gin.H{
		"message": "Аккаунт активирован",
	})
}

func (ac *AuthController) ResendActivationCode(c *gin.Context) {
	//Привет саня это твой тайный поклонник

	currentUser := securityService.GetCurrentUser(c)

	if currentUser.IsActivate {
		c.JSON(http.StatusOK, gin.H{
			"message": "Аккаунт уже активирован",
		})
		return
	}

	code := strconv.Itoa(util.GetRandomNumber())

	database.AllDatabases.DbCode.Db.Set(context.Background(), code, currentUser.Id, time.Minute*2)

	c.JSON(http.StatusOK, gin.H{
		"code": code,
	})
}
