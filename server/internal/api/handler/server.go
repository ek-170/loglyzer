package handler

import (
	"strings"

	"github.com/ek-170/loglyzer/internal/config"
	"github.com/labstack/echo-contrib/pprof"
	"github.com/labstack/echo/v4"
)

const (
	apiPathV1 = "/api/v1/"
)

func StartMainServer() {

	e := echo.New()
	pprof.Register(e)
	/* handle "hello" */
	e.GET(joinPathV1("hello"), HandleHello)

	/* handle "Grok" */
	e.POST(joinPathV1("grok-patterns"), HandleGrokFind)
	e.PUT(joinPathV1("grok-patterns/:grok-id"), HandleGrokCreate)
	e.DELETE(joinPathV1("grok-patterns/:grok-id"), HandleGrokDelete)

	/* handle "Analysis" */
	e.POST(joinPathV1("analyses"), HandleAnalysisFind)
	e.GET(joinPathV1("analyses/:analysis"), HandleAnalysisGet)
	e.PUT(joinPathV1("analyses/:analysis"), HandleAnalysisCreate)
	e.DELETE(joinPathV1("analyses/:analysis"), HandleAnalysisDelete)

	/* handle "ParseSource" */
	e.POST(joinPathV1("analyses/:analysis/parse-sources"), HandleParseSourceFind)
	e.POST(joinPathV1("analyses/:analysis/parse-sources/new"), HandleParseSourceCreate)
	e.DELETE(joinPathV1("analyses/:analysis/parse-sources/:parse-source-id"), HandleParseSourceDelete)

	/* handle "File" */
	e.POST(joinPathV1("files"), HandleFileFind)

	e.Logger.Fatal(e.Start(":" + config.Config.Server.Port))
	
}

func joinPathV1(pattern string) string {
	if strings.HasPrefix(pattern, "/") {
		return apiPathV1 + pattern[1:]
	}
	return apiPathV1 + pattern
}
