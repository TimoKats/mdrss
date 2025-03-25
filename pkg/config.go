// Submodule for reading config file and using the config object.

package lib

import (
	"errors"
	"os"

	"gopkg.in/ini.v1"
)

func fileExists(filename string) bool {
	if _, err := os.Stat(filename); err != nil {
		return false
	}
	return true
}

func (config Config) valid() bool {
	return len(config.OutputFolder) > 0 || len(config.OutputFile) > 0
}

func ReadConfig(filePath string) (Config, error) {
	var config Config
	if !fileExists(filePath) {
		return config, errors.New("config file not found. Please add it at ~/.mdrss")
	}
	fileContent, readErr := ini.Load(filePath)
	if readErr != nil {
		return config, readErr
	}
	parseErr := fileContent.MapTo(&config)
	if parseErr == nil && !config.valid() {
		parseErr = errors.New("please set OutputFile or OutputFolder in config")
	}
	return config, parseErr
}
