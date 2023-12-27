package controller

import (
	"github.com/MeibisuX673/LessonRedis/app/controller/apiController/authController"
	"github.com/MeibisuX673/LessonRedis/app/controller/apiController/userController"
)

type Controller struct {
	AuthController authController.AuthController
	UserController userController.UserController
}
