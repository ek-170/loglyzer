package main

import (
	"log"

	"github.com/ek-170/loglyzer/internal/api/handler"
)

func main() {
	logo := `
   _                _                       
  | |    ___   __ _| |   _   _ _______ _ __ 
  | |   / _ \ / _\` + "`" + `| |  | | | |_  / _ \ '__|
  | |__| (_) | (_| | |__| |_| |/ /  __/ |   
  |_____\___/ \__, |_____\__, /___\___|_|   
              |___/      |___/        
`
	log.Println(logo)
	handler.StartMainServer()
}