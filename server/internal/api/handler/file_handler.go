package handler

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/ek-170/loglyzer/internal/usecase"
)

func HandleFileFind(c echo.Context) error {
	log.Println("Start finding Files.")
	q := c.QueryParam("q")
	log.Printf("query keyword is \"%s\"", q)
	usecase := usecase.NewFileUsecase()
	res, err := usecase.FindFiles(q)
	if err != nil {
	  return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, res)
}