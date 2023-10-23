package config

import (
	"testing"
)

func TestLoadConfig(t *testing.T) {
  // check to open file and read contens.
  LoadConfig("config_test.yml")

  expectedServerPort := "19999"
  loadedServerPort := Config.Server.Port
  if loadedServerPort != expectedServerPort {
    t.Errorf("expected %v, but %v", expectedServerPort, loadedServerPort)
  }

  expectedFulTextSearchPort := "29999"
  loadedFulTextSearchPort := Config.FullTextSearch.Port
  if loadedFulTextSearchPort != expectedFulTextSearchPort {
    t.Errorf("expected %v, but %v", expectedFulTextSearchPort, loadedFulTextSearchPort)
  }
}
