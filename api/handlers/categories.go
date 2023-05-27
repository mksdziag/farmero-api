package handlers

import (
	"errors"
	"net/http"

	api "github.com/mksdziag/farmer-api/api/errors"

	"github.com/labstack/echo/v4"
	"github.com/mksdziag/farmer-api/features/categories"
)

func GetCategories(c echo.Context) error {
	categories, err := categories.GetCategories()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, api.CreateApiError(err.Error(), http.StatusInternalServerError))
	}
	return c.JSON(http.StatusOK, categories)
}

func GetCategory(c echo.Context) error {
	id := c.Param("id")

	category, err := categories.GetCategory(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, api.CreateApiError(err.Error(), http.StatusNotFound))
	}

	return c.JSON(http.StatusOK, category)
}

func CreateCategory(c echo.Context) error {
	categoryPayload := categories.Category{}

	err := c.Bind(&categoryPayload)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, api.CreateApiError(err.Error(), http.StatusInternalServerError))
	}

	if err := c.Validate(categoryPayload); err != nil {
		return c.JSON(http.StatusBadRequest, api.CreateApiError(err.Error(), http.StatusBadRequest))
	}

	createdCategory, err := categories.CreateCategory(categoryPayload)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, api.CreateApiError(err.Error(), http.StatusInternalServerError))
	}

	return c.JSON(http.StatusCreated, createdCategory)
}

func UpdateCategory(c echo.Context) error {
	id := c.Param("id")
	categoryPayload := categories.Category{}

	err := c.Bind(&categoryPayload)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, api.CreateApiError(err.Error(), http.StatusInternalServerError))
	}

	if err := c.Validate(categoryPayload); err != nil {
		return c.JSON(http.StatusBadRequest, api.CreateApiError(err.Error(), http.StatusBadRequest))
	}

	updatedCategory, err := categories.UpdateCategory(id, categoryPayload)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, api.CreateApiError(err.Error(), http.StatusInternalServerError))
	}

	return c.JSON(http.StatusOK, updatedCategory)
}

func DeleteCategory(c echo.Context) error {
	id := c.Param("id")

	err := categories.DeleteCategory(id)
	if err != nil {
		customErr := errors.New("item not found")
		return c.JSON(http.StatusInternalServerError, api.CreateApiError(customErr.Error(), http.StatusNotFound))
	}

	return c.NoContent(http.StatusOK)
}
