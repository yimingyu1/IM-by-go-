package utils

import (
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var (
	DB *gorm.DB
)

func InitConfig() error {
	viper.SetConfigName("app")
	viper.AddConfigPath("config")
	if err := viper.ReadInConfig(); err != nil {
		log.Println("read config err", err)
		return err
	}
	return nil
}

func InitDB() (err error) {
	DB, err = gorm.Open(mysql.Open(viper.GetString("mysql.dsn")), &gorm.Config{})
	if err != nil {
		return err
	}
	return nil
}
