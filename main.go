package main

import (
	"flag"
	"github.com/axzed/gohub/bootstrap"
	"github.com/axzed/gohub/pkg/config"
	"github.com/gin-gonic/gin"
)

func main() {

	// 配置初始化，依赖命令行 --env 参数
	var env string
	flag.StringVar(&env, "env", "", "加载 .env 文件，如 --env=testing 加载的是 .env.testing 文件")
	flag.Parse()
	config.InitConfig(env)

	// 初始化 Gin 实例
	router := gin.Default()

	// 初始化路由绑定
	bootstrap.SetupRoute(router)

	// 运行服务
	err := router.Run(":3000")
	if err != nil {
		panic(err)
	}

}
