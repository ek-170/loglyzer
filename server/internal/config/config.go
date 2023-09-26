package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Server struct {
  Port string
}

type ConfigList struct {
	Server Server
}

var Config ConfigList

func init() {
	LoadConfig()
}

func LoadConfig() {
  f, err := os.Open("../../config.yml")
  if err != nil {
    log.Fatal("config.yml loading error:", err)
  }
  defer f.Close()

  err = yaml.NewDecoder(f).Decode(&Config)
}
