package handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ek-170/loglyzer/internal/consts"
	"github.com/ek-170/loglyzer/internal/domain/repository"
	"github.com/ek-170/loglyzer/internal/usecase"
	"github.com/labstack/echo/v4"
)

func HandleSearchTargetFind(c echo.Context) error {
	log.Println("Start finding Search Target.")
	usecase := usecase.NewSearchTargetUsecase(repository.NewEsSearchTargetRepository())
	// only search for property "alias"
	q := c.QueryParam("q")
	log.Printf("query keyword is \"%s\"", q)
	st, err := usecase.FindSearchTargets(q)
	if err != nil {
	  echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf(consts.FAIL_FIND, "Search Target"))
	}
	return c.JSON(http.StatusOK, st)
}

func HandleSearchTargetGet(c echo.Context) error {
	log.Println("Start fetching Search Target.")
	usecase := usecase.NewSearchTargetUsecase(repository.NewEsSearchTargetRepository())
	name := c.Param("name")
	log.Printf("specified name is \"%s\"", name)
	st, err := usecase.GetSearchTarget(name)
	if err != nil {
	  echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf(consts.FAIL_GET, "Search Target"))
	}
	return c.JSON(http.StatusOK, st)
}

func HandleSearchTargetCreate(c echo.Context) error {
	log.Println("Start creating Search Target.")
	usecase := usecase.NewSearchTargetUsecase(repository.NewEsSearchTargetRepository())
	name := c.Param("name")
	log.Printf("specified name is \"%s\"", name)
	err := usecase.CreateSearchTarget(name)
	if err != nil {
	  echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf(consts.FAIL_CREATE, "Search Target"))
	}
	return c.NoContent(http.StatusCreated)
}

func HandleSearchTargetDelete(c echo.Context) error {
	log.Println("Start deleting Search Target.")
	usecase := usecase.NewSearchTargetUsecase(repository.NewEsSearchTargetRepository())
	name := c.Param("name")
	log.Printf("specified name is \"%s\"", name)
	err := usecase.DeleteSearchTarget(name)
	if err != nil {
	  echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf(consts.FAIL_DELETE, "Search Target"))
	}
	return c.NoContent(http.StatusOK)
}