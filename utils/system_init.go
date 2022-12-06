package utils

import (
	"gin_chat/common"
	"gin_chat/model"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
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
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: true,
			Colorful:                  true,
		},
	)
	common.DB, err = gorm.Open(mysql.Open(viper.GetString("mysql.dsn")), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		return err
	}
	err = common.DB.AutoMigrate(&model.UserBasic{})
	if err != nil {
		return err
	}
	return nil
}
