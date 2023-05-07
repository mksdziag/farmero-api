package api

import (
	"github.com/labstack/echo/v4"
	"github.com/mksdziag/farmer-api/api/handlers"
)

func StartServer() {
	app := echo.New()

	apiGroup := app.Group("/api")
	apiGroup.GET("/status", handlers.GetStatus)

	articlesGroup := apiGroup.Group("/articles")
	articlesGroup.GET("", handlers.GetArticles)
	articlesGroup.GET("/category/:categoryId", handlers.GetArticlesByCategory)
	articlesGroup.GET("/:id", handlers.GetArticle)
	articlesGroup.POST("/", handlers.CreateArticle)
	articlesGroup.PATCH("/:id", handlers.UpdateArticle)
	articlesGroup.DELETE("/:id", handlers.DeleteArticle)

	app.Logger.Fatal(app.Start(":5000"))
}
