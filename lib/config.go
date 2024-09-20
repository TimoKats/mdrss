package lib

import (
  "encoding/json"
  "errors"
  "os"
)

func FileExists(filename string) bool {
  if _, err := os.Stat(filename); err != nil {
    return false
  }
  return true
}

func ReadConfig(configPath string) (Config, error) {
  var config Config
  if FileExists(configPath) {
    configContent, _ := os.ReadFile(configPath)
    jsonCfgErr := json.Unmarshal(configContent, &config)
    if jsonCfgErr != nil {
      return config, errors.New("Error when reading config file.")
    }
    return config, nil
  }
  return config, errors.New("Config file not found. Please add it at ~/.mdrss/config.json or run \"mdrss init\"")
}
