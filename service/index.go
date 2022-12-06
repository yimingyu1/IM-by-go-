package service

import (
	"gin_chat/common"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetIndex
// @Summary gin_chat入口
// @Description index
// @Tags 首页
// @Produce json
//
//	@Success 200 {string} string "ok"
//
// @Router / [get]
func GetIndex(c *gin.Context) {
	c.JSON(http.StatusOK, common.BuildSuccessResponseNoData())
}
