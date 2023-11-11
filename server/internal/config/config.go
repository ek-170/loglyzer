package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Server struct {
  Port string `yaml:"port"`
  LogDir string `yaml:"logDir"`
}

type FullTextSearch struct {
  Schme string `yaml:"ftsScheme"`
  Host string  `yaml:"ftsHost"`
  Port string  `yaml:"ftsPort"`
}

type ConfigList struct {
	Server Server                 `yaml:"server"`
  FullTextSearch FullTextSearch `yaml:"fullTextSearch"`
}

var Config ConfigList

func init() {
	LoadConfig("config.yml")
}

func LoadConfig(fileName string) {
  // search from server dir
  f, err := os.Open("./" + fileName)
  if err != nil {
    // serarch from main.go
    f, err = os.Open("../../" + fileName)
      if err != nil {
        log.Fatal(fileName + " loading error:", err)
      }
  }
  defer f.Close()

  err = yaml.NewDecoder(f).Decode(&Config)
  if err != nil {
    log.Fatal("config.yml decoding error:", err)
  }
}