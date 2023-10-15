package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Server struct {
  Port string `yaml:"port"`
}

type FullTextSearch struct {
  Schme string `yaml:"ftsScheme"`
  Host string  `yaml:"ftsHost"`
  Port string  `yaml:"ftsPort"`
}

type ConfigList struct {
	Server Server                 `yaml:",inline"`
  FullTextSearch FullTextSearch `yaml:",inline"`
}

var Config ConfigList

func init() {
	LoadConfig()
}

func LoadConfig() {
  // search from server dir
  f, err := os.Open("./config.yml")
  if err != nil {
    // serarch from main.go
    f, err = os.Open("../../config.yml")
      if err != nil {
        log.Fatal("config.yml loading error:", err)
      }
  }
  defer f.Close()

  err = yaml.NewDecoder(f).Decode(&Config)
  if err != nil {
    log.Fatal("config.yml decoding error:", err)
  }
}