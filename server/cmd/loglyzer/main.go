package main

import (
	"log"

	"github.com/ek-170/loglyzer/internal/api/handler"
)

func main() {
  log.Printf("Start Loglyzer API")
  handler.StartMainServer()
}