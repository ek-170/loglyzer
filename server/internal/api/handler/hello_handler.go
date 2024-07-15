package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func HandleHello(c echo.Context) error {
  return c.String(http.StatusOK, "Hello, World!")
}