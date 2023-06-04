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
	articlesGroup.GET("/category/:categoryId", handlers.GetArticlesByCategoryId)
	articlesGroup.GET("/category/key/:categoryKey", handlers.GetArticlesByCategoryKey)
	articlesGroup.GET("/:id", handlers.GetArticle)
	articlesGroup.POST("", handlers.CreateArticle)
	articlesGroup.PATCH("/:id", handlers.UpdateArticle)
	articlesGroup.DELETE("/:id", handlers.DeleteArticle)

	categoriesGroup := apiGroup.Group("/categories")
	categoriesGroup.GET("", handlers.GetCategories)
	categoriesGroup.GET("/:id", handlers.GetCategory)
	categoriesGroup.GET("/key/:key", handlers.GetCategoryByKey)
	categoriesGroup.POST("", handlers.CreateCategory)
	categoriesGroup.PATCH("/:id", handlers.UpdateCategory)
	categoriesGroup.DELETE("/:id", handlers.DeleteCategory)

	tagsGroup := apiGroup.Group("/tags")
	tagsGroup.GET("", handlers.GetTags)
	tagsGroup.GET("/:id", handlers.GetTag)
	tagsGroup.POST("", handlers.CreateTag)
	tagsGroup.PATCH("/:id", handlers.UpdateTag)
	tagsGroup.DELETE("/:id", handlers.DeleteTag)

	app.Logger.Fatal(app.Start(":5000"))
}
