package model

import (
	"gin_chat/utils"
	"gorm.io/gorm"
	"strconv"
	"time"
)

type UserIdParam struct {
	Id uint `json:"id"`
}

type UserParam struct {
	UserIdParam
	Name       string `json:"name" binding:"required"`
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"rePassword" binding:"required"`
	Phone      string `json:"phone" binding:"required" valid:"matches(^1[3-9]{1}\\d{9}$)"`
	Email      string `json:"email" binding:"required" valid:"email"`
}

type UpdateUserPwdParam struct {
	UserIdParam
	Password    string `json:"password" binding:"required"`
	NewPassword string `json:"newPassword" binding:"required"`
}

type LoginParam struct {
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Param2Mode(userParam UserParam) *UserBasic {
	now := time.Now()
	salt := strconv.Itoa(int(utils.GetRandWithIn1000()))
	passWord := utils.EncodePwd(userParam.Password, salt)
	userBasic := &UserBasic{
		Model:         gorm.Model{ID: userParam.Id},
		Name:          userParam.Name,
		Salt:          salt,
		Password:      passWord,
		Phone:         userParam.Phone,
		Email:         userParam.Email,
		LoginTime:     now,
		HeartBeatTime: now,
		LoginOutTime:  now,
		isLogout:      false,
	}
	return userBasic
}
