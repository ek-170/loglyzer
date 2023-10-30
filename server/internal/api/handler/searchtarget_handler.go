package handler

import (
	"log"
	"net/http"

	"github.com/ek-170/loglyzer/internal/domain/repository"
	"github.com/ek-170/loglyzer/internal/usecase"
	"github.com/labstack/echo/v4"
)

func HandleSearchTargetFind(c echo.Context) error {
	log.Println("Start finding SearchTarget.")
	usecase := usecase.NewSearchTargetUsecase(repository.NewEsSearchTargetRepository())
	// only search for property "alias"
	q := c.QueryParam("q")
	log.Printf("query keyword is \"%s\"", q)
	st, err := usecase.FindSearchTargets(q)
	if err != nil {
	  return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, st)
}

func HandleSearchTargetGet(c echo.Context) error {
	log.Println("Start fetching SearchTarget.")
	usecase := usecase.NewSearchTargetUsecase(repository.NewEsSearchTargetRepository())
	name := c.Param("search-target")
	log.Printf("specified SearchTarget name is \"%s\"", name)
	st, err := usecase.GetSearchTarget(name)
	if err != nil {
	  return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, st)
}

func HandleSearchTargetCreate(c echo.Context) error {
	log.Println("Start creating SearchTarget.")
	usecase := usecase.NewSearchTargetUsecase(repository.NewEsSearchTargetRepository())
	name := c.Param("search-target")
	log.Printf("specified new SearchTarget name is \"%s\"", name)
	err := usecase.CreateSearchTarget(name)
	if err != nil {
	  return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusCreated)
}

func HandleSearchTargetDelete(c echo.Context) error {
	log.Println("Start deleting SearchTarget.")
	usecase := usecase.NewSearchTargetUsecase(repository.NewEsSearchTargetRepository())
	name := c.Param("name")
	log.Printf("specified SearchTarget name is \"%s\"", name)
	err := usecase.DeleteSearchTarget(name)
	if err != nil {
	  return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusOK)
}