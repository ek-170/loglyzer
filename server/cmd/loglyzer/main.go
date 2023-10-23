package main

import (
	"log"

	"github.com/ek-170/loglyzer/internal/api/handler"
)

const (
    apiPathV1 = "/api/v1/"
  )

func main() {
  log.Printf("Start Loglyzer API")
  handler.StartMainServer()
}