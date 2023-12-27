package router

import (
	"github.com/MeibisuX673/LessonRedis/app/controller"
	"github.com/MeibisuX673/LessonRedis/app/middleware"
	"github.com/gin-gonic/gin"
)

var controllers controller.Controller

func initApiRouter(ge *gin.Engine) {

	controllers = initializationController()

	groupApi := ge.Group("/api")

	initAuthRoutes(groupApi)
	initUserRouters(groupApi)

}

func initUserRouters(rg *gin.RouterGroup) {

	user := rg.Group("users")
	{
		user.GET("/:id", controllers.UserController.GetItem)
		user.PUT("/:id", middleware.JwtMiddleware, controllers.UserController.Update)
	}

}

func initAuthRoutes(rg *gin.RouterGroup) {

	auth := rg.Group("auth")
	{
		auth.POST("/signUp", controllers.AuthController.SignUp)
		auth.POST("/activate", middleware.JwtMiddleware, controllers.AuthController.ActivateUser)
		auth.POST("/resendActivateCode", middleware.JwtMiddleware, controllers.AuthController.ResendActivationCode)
	}
}

func initializationController() controller.Controller {

	return controller.Controller{}
}
