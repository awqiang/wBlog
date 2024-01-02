package routers

import (
	"github.com/awqiang/wBlog/handlers"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/articles", handlers.GetArticles)
	router.GET("/articles/:id", handlers.GetArticle)
	router.POST("/articles", handlers.AddArticle)
	router.PUT("/articles/:id", handlers.EditArticle)
	router.DELETE("/articles/:id", handlers.DeleteArticle)

	return router

}
