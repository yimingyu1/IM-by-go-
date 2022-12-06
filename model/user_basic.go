package model

import (
	"gin_chat/common"
	"gorm.io/gorm"
	"log"
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
	LoginOutTime  time.Time
	isLogout      bool
	DeviceInfo    string
}

// TableName 自定义表名 user_basic
func (user *UserBasic) TableName() string {
	return "user_basic"
}

func GetUserList() []*UserBasic {
	data := make([]*UserBasic, 10)
	common.DB.Find(&data)
	return data
}

func CreateUser(user *UserBasic) {
	common.DB.Debug().Create(user)
}

func QueryUserById(id int) *UserBasic {
	var user UserBasic
	common.DB.Debug().First(&user, "id = ?", id)
	log.Println(user)
	return &user
}

func DeleteUser(user *UserBasic) {
	common.DB.Delete(user)
}
