package lib

import (
  "encoding/json"
  "errors"
  "os"
)

func ReadConfig(filename string) (Config, error) {
	var config Config
  configContent, _ := os.ReadFile(filename)
	jsonErr := json.Unmarshal(configContent, &config)
	if jsonErr != nil {
		return config, errors.New("Error when reading config file.")
	}
	return config, nil
}

