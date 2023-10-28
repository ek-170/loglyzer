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
  e.GET(joinPath("hello"), HandleHello)

  /* handle "Grok" */
  e.POST(joinPath("grok-patterns"), HandleGrokFind)
  // e.PUT(joinPath("grok-patterns"), HandleGrokPut)
  // e.DELETE(joinPath("grok-patterns"), HandleGrokDelete)

  /* handle "Search Target" */
  e.POST(joinPath("search-targets"), HandleSearchTargetFind)
  e.GET(joinPath("search-targets/:name"), HandleSearchTargetGet)
  e.PUT(joinPath("search-targets/:name"), HandleSearchTargetCreate)
  e.DELETE(joinPath("search-targets/:name"), HandleSearchTargetDelete)

  /* handle "Parce Source" */
  // e.GET(joinPath("parse-sources"), HandleParseSourcekGet)
  // e.PUT(joinPath("parse-sources"), HandleParseSourceGet)
  // e.DELETE(joinPath("parse-sources"), HandleParseSourceGet)

  e.Logger.Fatal(e.Start(":"+config.Config.Server.Port))
}

func joinPath(pattern string) string {
  if strings.HasPrefix(pattern, "/"){
     return apiPathV1 + pattern[1:]
  }
  return apiPathV1 + pattern
}