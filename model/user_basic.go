package model

import (
	"gorm.io/gorm"
	"time"
)

// UserBasic 用户基础信息Model
type UserBasic struct {
	gorm.Model
	Name          string
	Password      string
	Phone         string
	Email         string
	Identity      string
	ClientIp      string
	LoginTime     time.Time
	HeartBeatTime time.Time
	LogOutTime    time.Time
	isLogout      bool
	DeviceInfo    string
}

// TableName 自定义表名 user_basic
func (user *UserBasic) TableName() string {
	return "user_basic"
}
