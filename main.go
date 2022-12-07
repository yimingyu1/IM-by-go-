package main

import (
	"gin_chat/init_boot"
	"gin_chat/router"
	"log"
)

// 入口文件
func main() {
	configErr := init_boot.InitConfig()
	if configErr != nil {
		log.Println("config err", configErr)
		return
	}
	dbErr := init_boot.InitDB()
	if dbErr != nil {
		log.Println("db err", dbErr)
		return
	}
	engine := router.Router()
	err := engine.Run(":9999")
	if err != nil {
		return
	}
}
