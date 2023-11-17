package handler

import (
	"log"
	"net/http"

	"github.com/ek-170/loglyzer/internal/domain/repository"
	"github.com/ek-170/loglyzer/internal/usecase"
	"github.com/labstack/echo/v4"
)

func HandleParseSourceFind(c echo.Context) error {
	log.Println("Start finding ParseSource.")
	usecase := usecase.NewParseSourceUsecase(repository.NewEsParseSourceRepository())
	q := c.QueryParam("q")
	log.Printf("query keyword is \"%s\"", q)
	st, err := usecase.FindParseSources(q)
	if err != nil {
	  return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, st)
}

func HandleParseSourceGet(c echo.Context) error {
	log.Println("Start fetching ParseSource.")
	usecase := usecase.NewParseSourceUsecase(repository.NewEsParseSourceRepository())
	name := c.Param("name")
	log.Printf("specified name is \"%s\"", name)
	st, err := usecase.GetParseSource(name)
	if err != nil {
	  return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, st)
}
type ParseSourceCreateRequest struct {
	File string     `json:"file"`
	MultiLine bool  `json:"multiLine"`
	GrokId string   `json:"grokId"`
}

func HandleParseSourceCreate(c echo.Context) error {
	log.Println("Start creating ParseSource.")
	req := ParseSourceCreateRequest{}
  if err := c.Bind(&req); err != nil {
    return c.JSON(http.StatusBadRequest, "bad request")
  }
	searchTarget := c.Param("search-target")
	log.Printf("parsing target file name is \"%s\"", req.File)
	log.Printf("MultiLine setting enabled is \"%t\"", req.MultiLine)
	log.Printf("Grok pattern name use for parsing is \"%s\"", req.GrokId)
	log.Printf("specified SearchTarget is \"%s\"", searchTarget)
	usecase := usecase.NewParseSourceUsecase(repository.NewEsParseSourceRepository())
	err := usecase.CreateParseSource(searchTarget, req.MultiLine, req.File, req.GrokId)
	if err != nil {
	  return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusCreated)
}

func HandleParseSourceDelete(c echo.Context) error {
	log.Println("Start deleting ParseSource.")
	usecase := usecase.NewParseSourceUsecase(repository.NewEsParseSourceRepository())
	name := c.Param("name")
	log.Printf("specified name is \"%s\"", name)
	err := usecase.DeleteParseSource(name)
	if err != nil {
	  return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusOK)
}