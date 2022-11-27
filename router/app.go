package router

import (
	"gin_chat/service"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	engine := gin.Default()
	engine.GET("/", service.GetIndex)
	return engine
}
