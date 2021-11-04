package handlers

import (
	"github.com/gin-gonic/gin"
	"mdblog/utils"
	"net/http"
	"strconv"
)

// GetAllArticle 获取全部文章
func GetAllArticle(ctx *gin.Context) {
	articles, err := utils.GetMarkdowns()
	// 未能成功获取文章列表
	if err != nil {
		InternalError := gin.H{
			"Error": "can't get article list!",
			"Code":  "500",
		}
		ctx.JSON(http.StatusInternalServerError, InternalError)
		return
	}
	ctx.JSON(http.StatusOK, articles)
}

// GetArticleByID 按照文章 ID 获取文章
func GetArticleByID(ctx *gin.Context) {
	ID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		NotFoundError := gin.H{
			"Error": "Can't find article",
			"Code":  "404",
		}
		ctx.JSON(http.StatusNotFound, NotFoundError)
		return
	}

	article, err := utils.GetMarkdownByID(ID)

	// 未找到文章
	if err != nil {
		NotFoundError := gin.H{
			"Error": "can't find the article",
			"Code":  "404",
		}
		ctx.JSON(http.StatusNotFound, NotFoundError)
	}

	ctx.JSON(http.StatusOK, article)
}
