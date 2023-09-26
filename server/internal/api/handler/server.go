package handler

import (
	"net/http"

	"github.com/ek-170/loglyzer/internal/config"
)

const (
  apiPathV1 = "/api/v1/"
)

func StartMainServer() error {
	http.HandleFunc(apiPathV1+"hello", NewHelloHandler().HandleHello)
	// http.HandleFunc(apiPathV1+"searchsource/", dh.HandleDiary)

	return http.ListenAndServe(":"+config.Config.Server.Port, nil)
}
