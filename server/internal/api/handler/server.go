package handler

import (
	"strings"

	"github.com/ek-170/loglyzer/internal/config"
	"github.com/labstack/echo/v4"
)

const (
  apiPathV1 = "/api/v1/"
)

func StartMainServer() {

  e := echo.New()

  /* handle "hello" */
  e.GET(joinPathV1("hello"), HandleHello)

  /* handle "Grok" */
  e.POST(joinPathV1("grok-patterns"), HandleGrokFind)
  // e.PUT(joinPathV1("grok-patterns"), HandleGrokPut)
  // e.DELETE(joinPathV1("grok-patterns"), HandleGrokDelete)

  /* handle "SearchTarget" */
  e.POST(joinPathV1("search-targets"), HandleSearchTargetFind)
  e.GET(joinPathV1("search-targets/:search-target"), HandleSearchTargetGet)
  e.PUT(joinPathV1("search-targets/:search-target"), HandleSearchTargetCreate)
  e.DELETE(joinPathV1("search-targets/:search-target"), HandleSearchTargetDelete)

  /* handle "ParseSource" */
  // e.POST(joinPathV1("search-targets/:search-targets-name/parse-sources/:parse-sources-name"), HandleParseSourceFind)
  // e.GET(joinPathV1("search-targets/:search-targets-name/parse-sources/:parse-sources-name"), HandleParseSourceGet)
  e.POST(joinPathV1("search-targets/:search-target/parse-sources"), HandleParseSourceCreate)
  // e.DELETE(joinPathV1("search-targets/:search-targets-name/parse-sources/:parse-sources-name"), HandleParseSourceDelete)

  /* handle "File" */
  e.POST(joinPathV1("files"), HandleFileFind)

  e.Logger.Fatal(e.Start(":"+config.Config.Server.Port))
}

func joinPathV1(pattern string) string {
  if strings.HasPrefix(pattern, "/"){
     return apiPathV1 + pattern[1:]
  }
  return apiPathV1 + pattern
}