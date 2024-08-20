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

func getConfigPath(configPathFlag string) string {
  if len(configPathFlag) > 0 {
    Warn.Printf("Using config path from flag: %s", configPathFlag)
    return configPathFlag
  } else {
    dirname, _ := os.UserHomeDir()
    return dirname + "/.mdrss/config.json"
  }
}

func ReadConfig(configPathFlag string) (Config, error) {
  var config Config
  configPath := getConfigPath(configPathFlag)
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

