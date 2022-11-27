package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetIndex(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"code":    0,
		"data":    nil,
	})
}
