package main

import (
	"gin_chat/router"
	"gin_chat/utils"
	"log"
)

// 入口文件
func main() {
	configErr := utils.InitConfig()
	if configErr != nil {
		log.Println("config err", configErr)
		return
	}
	dbErr := utils.InitDB()
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
