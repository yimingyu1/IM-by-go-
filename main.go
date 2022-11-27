package main

import "gin_chat/router"

// 入口文件
func main() {
	engine := router.Router()
	err := engine.Run(":9999")
	if err != nil {
		return
	}
}
