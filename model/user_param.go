package model

import (
	"gorm.io/gorm"
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

func Param2Mode(userParam UserParam) *UserBasic {
	now := time.Now()
	return &UserBasic{
		Model:         gorm.Model{ID: userParam.Id},
		Name:          userParam.Name,
		Password:      userParam.Password,
		Phone:         userParam.Phone,
		Email:         userParam.Email,
		LoginTime:     now,
		HeartBeatTime: now,
		LoginOutTime:  now,
		isLogout:      false,
	}
}
