package main

import (
	"github.com/gin-gonic/gin"
	"mdblog/middlewares"
	"mdblog/routers"
)

func main() {
	g := gin.Default()

	// 使用跨域中间件
	g.Use(middlewares.CORS())

	// 静态图片目录
	g.Static("/pics", "./pics")

	// 初始化路由
	routers.InitializeRouters(g)

	g.Run(":9999")
}
