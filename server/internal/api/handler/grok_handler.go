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

type GrokCreateRequest struct {
	Pattern string  `json:"pattern"`
	PatternDefs map[string]string `json:"pattern_definitions"`
  Description string `json:"description"`
}

func HandleGrokCreate(c echo.Context) error {
	log.Println("Start creating Grok Pattern.")
	req := GrokCreateRequest{}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, "bad request")
	}
	grokId := c.Param("grok-id")
	log.Printf("Grok Pattern is \"%s\"", req.Pattern)
	usecase := usecase.NewGrokUsecase(repository.NewEsGrokRepository())
	err := usecase.CreateGrokPatterns(grokId, req.Pattern, req.PatternDefs, req.Description)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusCreated)
}