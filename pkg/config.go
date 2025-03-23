// Submodule for reading config file and using the config object.

package lib

import (
  "gopkg.in/ini.v1"
  "errors"
)

func (config Config) valid() bool {
  return isSet(config.OutputFolder) || isSet(config.OutputFile)
}

func ReadConfig(filePath string) (Config, error) {
  var config Config
  if !fileExists(filePath) {
    return config, errors.New("Config file not found. Please add it at ~/.mdrss")
  }
  fileContent, readErr := ini.Load(filePath)
  if readErr != nil { return config, readErr }
  parseErr := fileContent.MapTo(&config)
  if parseErr == nil && !config.valid() {
    parseErr = errors.New("Please set OutputFile or OutputFolder in config.")
  }
  return config, parseErr
}

