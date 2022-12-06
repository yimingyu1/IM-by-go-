package service

import (
	"gin_chat/common"
	"gin_chat/model"
	"github.com/gin-gonic/gin"
	"log"
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

func CreateUser(c *gin.Context) {
	var userParam model.UserParam
	if c.ShouldBindJSON(&userParam) == nil {
		log.Println("createUser param is ", userParam)
		if userParam.Password != userParam.RePassword {
			c.JSON(http.StatusOK, common.BuildFailResponse("两次密码不一致"))
		} else {
			model.CreateUser(model.Param2Mode(userParam))
			c.JSON(http.StatusOK, common.BuildSuccessResponse("创建用户成功"))
		}
	} else {
		c.JSON(http.StatusOK, common.BuildFailResponse("系统错误"))
	}
}

func DeleteUser(c *gin.Context) {
	var userIdParam model.UserIdParam
	err := c.ShouldBindJSON(&userIdParam)
	if err == nil {
		user := model.QueryUserById(userIdParam.Id)
		if user.ID == 0 {
			c.JSON(http.StatusOK, common.BuildFailResponse("要删除的用户不存在"))
			return
		}
		model.DeleteUser(user)
		c.JSON(http.StatusOK, common.BuildSuccessResponse("删除成功"))
	} else {
		c.JSON(http.StatusOK, common.BuildFailResponse("请选择要删除的用户"))
	}
}
