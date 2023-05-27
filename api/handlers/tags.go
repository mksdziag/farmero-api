package handlers

import (
	"errors"
	"net/http"

	api "github.com/mksdziag/farmer-api/api/errors"

	"github.com/labstack/echo/v4"
	"github.com/mksdziag/farmer-api/features/tags"
)

func GetTags(c echo.Context) error {
	tags, err := tags.GetTags()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, api.CreateApiError(err.Error(), http.StatusInternalServerError))
	}
	return c.JSON(http.StatusOK, tags)
}

func GetTag(c echo.Context) error {
	id := c.Param("id")

	tag, err := tags.GetTag(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, api.CreateApiError(err.Error(), http.StatusNotFound))
	}

	return c.JSON(http.StatusOK, tag)
}

func CreateTag(c echo.Context) error {
	tagPayload := tags.Tag{}

	err := c.Bind(&tagPayload)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, api.CreateApiError(err.Error(), http.StatusInternalServerError))
	}

	if err := c.Validate(tagPayload); err != nil {
		return c.JSON(http.StatusBadRequest, api.CreateApiError(err.Error(), http.StatusBadRequest))
	}

	createdTag, err := tags.CreateTag(tagPayload)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, api.CreateApiError(err.Error(), http.StatusInternalServerError))
	}

	return c.JSON(http.StatusCreated, createdTag)
}

func UpdateTag(c echo.Context) error {
	id := c.Param("id")
	tagPayload := tags.Tag{}

	err := c.Bind(&tagPayload)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, api.CreateApiError(err.Error(), http.StatusInternalServerError))
	}

	if err := c.Validate(tagPayload); err != nil {
		return c.JSON(http.StatusBadRequest, api.CreateApiError(err.Error(), http.StatusBadRequest))
	}

	updatedTag, err := tags.UpdateTag(id, tagPayload)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, api.CreateApiError(err.Error(), http.StatusInternalServerError))
	}

	return c.JSON(http.StatusOK, updatedTag)
}

func DeleteTag(c echo.Context) error {
	id := c.Param("id")

	err := tags.DeleteTag(id)
	if err != nil {
		customErr := errors.New("item not found")
		return c.JSON(http.StatusInternalServerError, api.CreateApiError(customErr.Error(), http.StatusNotFound))
	}

	return c.NoContent(http.StatusOK)
}
