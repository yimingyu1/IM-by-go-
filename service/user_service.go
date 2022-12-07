package service

import (
	"gin_chat/common"
	"gin_chat/model"
	"github.com/asaskevich/govalidator"
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
		} else if _, err := govalidator.ValidateStruct(&userParam); err != nil {
			log.Println(err)
			Errors, ok := err.(govalidator.Errors)
			if ok {
				for _, errorItem := range Errors {
					goValidatorErr := errorItem.(govalidator.Error)
					if goValidatorErr.Name == "phone" {
						c.JSON(http.StatusOK, common.BuildFailResponse("手机号格式不正确"))
						return
					}
					if goValidatorErr.Name == "email" {
						c.JSON(http.StatusOK, common.BuildFailResponse("邮箱格式不正确"))
						return
					}
				}
			} else {
				c.JSON(http.StatusOK, common.BuildFailResponse("系统异常"))
			}
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

func UpdateUser(c *gin.Context) {
	var userParam model.UserParam
	err := c.ShouldBindJSON(&userParam)
	if err == nil {
		user := model.QueryUserById(userParam.Id)
		if user.ID == 0 {
			c.JSON(http.StatusOK, common.BuildFailResponse("要更新的用户不存在"))
			return
		}
		model.UpdateUser(model.Param2Mode(userParam))
		c.JSON(http.StatusOK, common.BuildSuccessResponse("更新成功"))
	} else {
		c.JSON(http.StatusOK, common.BuildFailResponse("请选择要修改的用户"))
	}
}
