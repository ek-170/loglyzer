package handler

import (
	"log"
	"net/http"

	fr "github.com/ek-170/loglyzer/internal/domain/filereader"
	"github.com/ek-170/loglyzer/internal/domain/repository"
	"github.com/ek-170/loglyzer/internal/usecase"
	"github.com/labstack/echo/v4"
)

func HandleParseSourceFind(c echo.Context) error {
	log.Println("Start finding ParseSource.")
	usecase := usecase.NewParseSourceUsecase(repository.NewEsParseSourceRepository())
	q := c.QueryParam("q")
	log.Printf("query keyword is \"%s\"", q)
	searchTarget := c.Param("search-target")
	log.Printf("specified SearchTarget is \"%s\"", searchTarget)
	st, err := usecase.FindParseSources(q, searchTarget)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, st)
}

type ParseSourceCreateRequest struct {
	FileReadMode string `json:"fileReadMode"`
	SshKeyPath   string `json:"sshKeyPath"`
	UserName     string `json:"userName"`
	Password     string `json:"password"`
	Host         string `json:"host"`
	Port         int    `json:"port"`
	FilePath     string `json:"filePath"`
	MultiLine    bool   `json:"multiLine"`
	GrokId       string `json:"grokId"`
}

func HandleParseSourceCreate(c echo.Context) error {
	log.Println("Start creating ParseSource.")
	req := ParseSourceCreateRequest{}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, "bad request")
	}
	searchTarget := c.Param("search-target")
	log.Printf("specified SearchTarget is \"%s\"", searchTarget)
	log.Printf("parsing target file path is \"%s\"", req.FilePath)
	log.Printf("MultiLine setting enabled is \"%t\"", req.MultiLine)
	log.Printf("Grok pattern name use for parsing is \"%s\"", req.GrokId)
	frConf := fr.FileReaderConfig{
		FileReadMode: req.FileReadMode,
		Path: req.FilePath,
		SshKeyPath: req.SshKeyPath,
		UserName: req.UserName,
		Password: req.Password,
		Host: req.Host,
		Port: &req.Port,
	}
	usecase := usecase.NewParseSourceUsecase(repository.NewEsParseSourceRepository())
	err := usecase.CreateParseSource(searchTarget, req.MultiLine, frConf, req.GrokId)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusCreated)
}

func HandleParseSourceDelete(c echo.Context) error {
	log.Println("Start deleting ParseSource.")
	usecase := usecase.NewParseSourceUsecase(repository.NewEsParseSourceRepository())
	searchTarget := c.Param("search-target")
	log.Printf("specified SearchTarget is \"%s\"", searchTarget)
	id := c.Param("parse-source-id")
	log.Printf("specified ParseSource info ID is \"%s\"", id)
	err := usecase.DeleteParseSource(id, searchTarget)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusOK)
}
