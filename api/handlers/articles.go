package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mksdziag/farmer-api/features/articles"
)

func GetArticles(c echo.Context) error {
	articles, err := articles.GetArticles()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, articles)
}

func GetArticle(c echo.Context) error {
	id := c.Param("id")

	article, err := articles.GetArticle(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, article)
}

func CreateArticle(c echo.Context) error {
	resData := make(map[string]string)
	resData["status"] = "OK"

	return c.JSON(http.StatusOK, resData)
}

func UpdateArticle(c echo.Context) error {
	resData := make(map[string]string)
	resData["status"] = "OK"

	return c.JSON(http.StatusOK, resData)
}

func GetArticlesByCategory(c echo.Context) error {
	category := c.Param("categoryId")

	articles, err := articles.GetArticlesByCategory(category)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, articles)
}

func DeleteArticle(c echo.Context) error {
	resData := make(map[string]string)
	resData["status"] = "OK"

	return c.JSON(http.StatusOK, resData)
}
