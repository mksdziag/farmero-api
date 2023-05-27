package api

import (
	"github.com/labstack/echo/v4"
	"github.com/mksdziag/farmer-api/api/handlers"
	"github.com/mksdziag/farmer-api/validator"
)

func StartServer() {
	app := echo.New()
	app.Validator = validator.NewAppValidator()

	apiGroup := app.Group("/api")
	apiGroup.GET("/status", handlers.GetStatus)

	articlesGroup := apiGroup.Group("/articles")
	articlesGroup.GET("", handlers.GetArticles)
	articlesGroup.GET("/category/:categoryId", handlers.GetArticlesByCategory)
	articlesGroup.GET("/:id", handlers.GetArticle)
	articlesGroup.POST("", handlers.CreateArticle)
	articlesGroup.PATCH("/:id", handlers.UpdateArticle)
	articlesGroup.DELETE("/:id", handlers.DeleteArticle)

	categoriesGroup := apiGroup.Group("/categories")
	categoriesGroup.GET("", handlers.GetCategories)
	categoriesGroup.GET("/:id", handlers.GetCategory)
	categoriesGroup.POST("", handlers.CreateCategory)
	categoriesGroup.PATCH("/:id", handlers.UpdateCategory)
	categoriesGroup.DELETE("/:id", handlers.DeleteCategory)

	app.Logger.Fatal(app.Start(":5000"))
}
