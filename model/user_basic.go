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
	Salt          string
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

func QueryUserById(id uint) *UserBasic {
	var user UserBasic
	common.DB.Debug().First(&user, "id = ?", id)
	log.Println(user)
	return &user
}

func QueryUserByName(name string) *UserBasic {
	var user UserBasic
	common.DB.Debug().First(&user, "name = ?", name)
	log.Println(user)
	return &user
}

func QueryUserByPhone(phone string) *UserBasic {
	var user UserBasic
	common.DB.Debug().First(&user, "phone = ?", phone)
	log.Println(user)
	return &user
}

func QueryUserByEmail(email string) *UserBasic {
	var user UserBasic
	common.DB.Debug().First(&user, "email = ?", email)
	log.Println(user)
	return &user
}

func DeleteUser(user *UserBasic) {
	common.DB.Debug().Delete(user)
}

func UpdateUser(user *UserBasic) {
	common.DB.Debug().Model(user).Updates(UserBasic{Name: user.Name, Phone: user.Phone, Email: user.Email})
}

func UpdateUserIdentity(id uint, identity string) {
	common.DB.Debug().Model(&UserBasic{}).Where("id = ?", id).Update("identity", identity)
}

func UpdateUserPwd(user *UserBasic) {
	common.DB.Model(user).Updates(UserBasic{Password: user.Password, Salt: user.Salt})
}
