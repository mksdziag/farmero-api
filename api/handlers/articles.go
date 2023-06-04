package handlers

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	api "github.com/mksdziag/farmer-api/api/errors"
	"github.com/mksdziag/farmer-api/features/articles"
	"github.com/mksdziag/farmer-api/features/categories"
	"github.com/mksdziag/farmer-api/features/tags"
)

func GetArticles(c echo.Context) error {
	articles, err := articles.GetArticles()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, api.CreateApiError(err.Error(), http.StatusInternalServerError))
	}

	for i, article := range articles {
		articles[i].Categories, err = categories.GetCategoriesByArticle(article.ID.String())
		if err != nil {
			return c.JSON(http.StatusInternalServerError, api.CreateApiError(err.Error(), http.StatusInternalServerError))
		}

		articles[i].Tags, err = tags.GetTagsByArticle(article.ID.String())
		if err != nil {
			return c.JSON(http.StatusInternalServerError, api.CreateApiError(err.Error(), http.StatusInternalServerError))
		}
	}

	return c.JSON(http.StatusOK, articles)
}

func GetArticle(c echo.Context) error {
	id := c.Param("id")

	article, err := articles.GetArticle(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, api.CreateApiError(err.Error(), http.StatusInternalServerError))
	}

	return c.JSON(http.StatusOK, article)
}

func CreateArticle(c echo.Context) error {
	articlePayload := articles.Article{}

	err := c.Bind(&articlePayload)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, api.CreateApiError(err.Error(), http.StatusInternalServerError))
	}

	if err := c.Validate(articlePayload); err != nil {
		return c.JSON(http.StatusBadRequest, api.CreateApiError(err.Error(), http.StatusBadRequest))
	}

	createdArticle, err := articles.CreateArticle(articlePayload)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, api.CreateApiError(err.Error(), http.StatusInternalServerError))
	}

	return c.JSON(http.StatusCreated, createdArticle)
}

func UpdateArticle(c echo.Context) error {
	id := c.Param("id")
	articlePayload := articles.Article{}

	err := c.Bind(&articlePayload)
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, api.CreateApiError(err.Error(), http.StatusInternalServerError))
	}

	if err := c.Validate(articlePayload); err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusBadRequest, api.CreateApiError(err.Error(), http.StatusBadRequest))
	}

	updatedArticle, err := articles.UpdateArticle(id, articlePayload)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, api.CreateApiError(err.Error(), http.StatusInternalServerError))
	}

	return c.JSON(http.StatusOK, updatedArticle)
}

func GetArticlesByCategoryId(c echo.Context) error {
	category := c.Param("categoryId")

	articles, err := articles.GetArticlesByCategoryId(category)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, api.CreateApiError(err.Error(), http.StatusInternalServerError))
	}

	return c.JSON(http.StatusOK, articles)
}
func GetArticlesByCategoryKey(c echo.Context) error {
	categoryKey := c.Param("categoryKey")

	articles, err := articles.GetArticlesByCategoryKey(categoryKey)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, api.CreateApiError(err.Error(), http.StatusInternalServerError))
	}

	return c.JSON(http.StatusOK, articles)
}

func DeleteArticle(c echo.Context) error {
	id := c.Param("id")

	err := articles.DeleteArticle(id)
	if err != nil {
		customErr := errors.New("item not found")
		return c.JSON(http.StatusInternalServerError, api.CreateApiError(customErr.Error(), http.StatusInternalServerError))
	}

	return c.NoContent(http.StatusOK)
}
