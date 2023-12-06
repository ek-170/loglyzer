package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Path struct {
  Base string
  Patterns string `yaml:"patterns"`
}

type Parser struct {
  Worker int          `yaml:"worker"`
  MultilineWorker int `yaml:"multilineWorker"`
}

type Server struct {
  Port string   `yaml:"port"`
  LogDir string `yaml:"logDir"`
}

type FullTextSearch struct {
  Schme string `yaml:"ftsScheme"`
  Host string  `yaml:"ftsHost"`
  Port string  `yaml:"ftsPort"`
  BulkUnit int `yaml:"ftsBulkUnit"`
}

type ConfigList struct {
  Path Path                     `yaml:"path"`
  Parser Parser                 `yaml:"parser"`
	Server Server                 `yaml:"server"`
  FullTextSearch FullTextSearch `yaml:"fullTextSearch"`
}

var Config ConfigList

func init() {
	LoadConfig("config.yml")
  setBasePath()
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

const EnvBasePathConfig = "BASE_PATH"

func setBasePath(){
  Config.Path.Base = os.Getenv(EnvBasePathConfig)
}