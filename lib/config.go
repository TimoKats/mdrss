package lib

import (
	"encoding/json"
	"errors"
	"os"
)

func getConfigPath() string {
  dirname, _ := os.UserHomeDir()
  return dirname + "/.mdrss/config.json"
}

func FileExists(filename string) bool {
  if _, err := os.Stat(filename); err != nil {
    return false
  }
  return true
}

func ReadConfig() (Config, error) {
  var config Config
  configPath := getConfigPath()
  if FileExists(configPath) {
    configContent, _ := os.ReadFile(configPath)
    jsonErr := json.Unmarshal(configContent, &config)
    if jsonErr != nil {
      return config, errors.New("Error when reading config file.")
    }
    return config, nil
  }
  return config, errors.New("Config file not found.")
}

