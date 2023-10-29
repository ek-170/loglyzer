package handler

import (
	"log"
	"net/http"

	"github.com/ek-170/loglyzer/internal/domain/repository"
	"github.com/ek-170/loglyzer/internal/usecase"
	"github.com/labstack/echo/v4"
)

func HandleGrokFind(c echo.Context) error {
  log.Println("Start finding Grok Patterns.")
  usecase := usecase.NewGrokUsecase(repository.NewEsGrokRepository())
  // only search for grok pattern name
  q := c.QueryParam("q")
  log.Printf("query keyword is \"%s\"", q)
  grok, err := usecase.FindGrokPatterns(q)
  if err != nil {
    return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
  }
  return  c.JSON(http.StatusOK, grok)
}