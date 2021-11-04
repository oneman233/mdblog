package routers

import (
	"github.com/gin-gonic/gin"
	"mdblog/handlers"
)

func InitializeRouters(g *gin.Engine) {
	// 路由组配置
	v1 := g.Group("/v1")
	{
		v1.GET("/articles", handlers.GetAllArticle)
		v1.GET("/article/:id", handlers.GetArticleByID)
	}
}
