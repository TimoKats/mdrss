package lib

import (
	"encoding/json"
	"errors"
	"os"
)

func getConfigPath(config_path string) string {
  if FileExists(config_path) {
    return config_path
  } else {
    dirname, _ := os.UserHomeDir()
    return dirname + "/.mdrss/config.json"
  }
}

func FileExists(filename string) bool {
  if _, err := os.Stat(filename); err != nil {
    return false
  }
  return true
}

func ReadConfig(config_path string) (Config, error) {
  var config Config
  configPath := getConfigPath(config_path)
  if FileExists(configPath) {
    configContent, _ := os.ReadFile(configPath)
    jsonErr := json.Unmarshal(configContent, &config)
    if jsonErr != nil {
      return config, errors.New("Error when reading config file.")
    }
    return config, nil
  }
  return config, errors.New("Config file not found. Please add it at ~/.mdrss/config.json")
}

