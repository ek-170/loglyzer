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

  // handle "hello"
  e.GET(joinPath("hello"), HandleHello)

  // handle "Grok"
  e.GET(joinPath("grok-patterns"), HandleGrokGet)

  e.Logger.Fatal(e.Start(":"+config.Config.Server.Port))
}

func joinPath(pattern string) string {
  if strings.HasPrefix(pattern, "/"){
     return apiPathV1 + pattern[1:]
  }
  return apiPathV1 + pattern
}