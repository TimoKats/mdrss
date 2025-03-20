package lib

import (
  "gopkg.in/ini.v1"
  "errors"
  "flag"
  "os"
)

func getCommand(arguments []string) (string, error) {
  validCommands := []string{"update", "ls", "conf", "init"}
  for _, validCommand := range validCommands {
    for _, argument := range arguments {
      if argument == validCommand {
        return validCommand, nil
      }
    }
  }
  return "", errors.New("No valid command found. Use mdrss <<ls, update, conf, init>>")
}

func defaultConfigPath() string {
  dirname, _ := os.UserHomeDir()
  return dirname + "/.mdrss"
}

func fileExists(filename string) bool {
  if _, err := os.Stat(filename); err != nil {
    return false
  }
  return true
}

func ParseArguments(arguments []string) (map[string]*string, error) {
  parsedArguments := make(map[string]*string)
  command, commandErr := getCommand(arguments)
  parsedArguments["config"] = flag.String("config", defaultConfigPath(), "path to config")
  parsedArguments["command"] = &command
  flag.Parse()
  return parsedArguments, commandErr
}

func ReadConfig(filePath string) (Config, error) {
  var config Config
  if !fileExists(filePath) {
    return config, errors.New("Feed file not found. Please add it at ~/.mdrss")
  }
  fileContent, readErr := ini.Load(filePath)
  if readErr != nil { return config, readErr }
  parseErr := fileContent.MapTo(&config)
  return config, parseErr
}

