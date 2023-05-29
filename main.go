package main

import (
	"github.com/axzed/gohub/bootstrap"
	"github.com/gin-gonic/gin"
)

func main() {

	// 初始化 Gin 实例
	router := gin.Default()

	// 初始化路由绑定
	bootstrap.SetupRoute(router)

	// 运行服务
	err := router.Run(":8000")
	if err != nil {
		panic(err)
	}

}
