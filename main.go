package main

import (
	"orderSystem/models"
	"orderSystem/routes"
)

func main() {
	// 引用数据库
	models.InitDb()
	// 引入路由组件
	routes.InitRouter()
}
