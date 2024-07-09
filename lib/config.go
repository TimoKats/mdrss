package lib

import (
	"encoding/json"
	"errors"
	"os"
)

func ReadConfig(configPath string) (Config, error) {
  var config Config
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

