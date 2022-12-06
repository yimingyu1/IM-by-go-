package service

import (
	"gin_chat/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetUserList(c *gin.Context) {
	data := model.GetUserList()
	c.JSON(http.StatusOK, gin.H{
		"success": data,
		"code":    0,
		"data":    nil,
	})
}
