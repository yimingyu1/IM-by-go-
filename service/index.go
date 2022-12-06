package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetIndex
// @BasePath /api/v1
// PingExample godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags 首页
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router / [get]
func GetIndex(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"code":    0,
		"data":    nil,
	})
}
