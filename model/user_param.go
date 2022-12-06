package model

import "time"

type UserIdParam struct {
	Id int `json:"id"`
}

type UserParam struct {
	UserIdParam
	Name       string `json:"name" binding:"required"`
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"rePassword" binding:"required"`
	Phone      string `json:"phone" binding:"required"`
	Email      string `json:"email" binding:"required"`
}

func Param2Mode(userParam UserParam) *UserBasic {
	now := time.Now()
	return &UserBasic{
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
