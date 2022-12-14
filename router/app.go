package router

import (
	"gin_chat/docs"
	"gin_chat/service"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Router() *gin.Engine {
	engine := gin.Default()
	docs.SwaggerInfo.BasePath = ""
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	engine.GET("/", service.GetIndex)
	userGroup := engine.Group("/user")
	{
		userGroup.POST("/createUser", service.CreateUser)
		userGroup.POST("/deleteUser", service.DeleteUser)
		userGroup.POST("/updateUser", service.UpdateUser)
		userGroup.POST("/updateUserPwd", service.UpdateUserPwd)
		userGroup.POST("/login", service.Login)
		userGroup.GET("/getUserList", service.GetUserList)
	}
	return engine
}
