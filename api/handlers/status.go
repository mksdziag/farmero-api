package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetStatus(c echo.Context) error {
	resData := make(map[string]string)
	resData["status"] = "OK"
	return c.JSON(http.StatusOK, resData)
}
