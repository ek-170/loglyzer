package handler

import (
	"log"
	"net/http"

	"github.com/ek-170/loglyzer/internal/domain/repository"
	"github.com/ek-170/loglyzer/internal/usecase"
	"github.com/labstack/echo/v4"
)

func HandleAnalysisFind(c echo.Context) error {
	log.Println("Start finding Analysis.")
	usecase := usecase.NewAnalysisUsecase(repository.NewEsAnalysisRepository())
	// only search for property "alias"
	q := c.QueryParam("q")
	log.Printf("query keyword is \"%s\"", q)
	st, err := usecase.FindAnalysiss(q)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, st)
}

func HandleAnalysisGet(c echo.Context) error {
	log.Println("Start fetching Analysis.")
	usecase := usecase.NewAnalysisUsecase(repository.NewEsAnalysisRepository())
	name := c.Param("analysis")
	log.Printf("specified Analysis name is \"%s\"", name)
	st, err := usecase.GetAnalysis(name)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, st)
}

func HandleAnalysisCreate(c echo.Context) error {
	log.Println("Start creating Analysis.")
	usecase := usecase.NewAnalysisUsecase(repository.NewEsAnalysisRepository())
	name := c.Param("analysis")
	log.Printf("specified new Analysis name is \"%s\"", name)
	err := usecase.CreateAnalysis(name)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusCreated)
}

func HandleAnalysisDelete(c echo.Context) error {
	log.Println("Start deleting Analysis.")
	usecase := usecase.NewAnalysisUsecase(repository.NewEsAnalysisRepository())
	name := c.Param("analysis")
	log.Printf("specified Analysis name is \"%s\"", name)
	err := usecase.DeleteAnalysis(name)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusOK)
}
