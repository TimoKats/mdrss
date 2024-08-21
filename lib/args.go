package lib

import (
  "errors"
  "flag"
  "os"
)

func getCommand(arguments []string) (string, error) {
  validCommands := []string{"update", "ls", "conf"}
  for _, validCommand := range validCommands {
    for _, argument := range arguments {
      if argument == validCommand {
        return validCommand, nil
      }
    }
  }
  return "", errors.New("No valid command found. Use mdrss <<ls, update, conf>>")
}

func defaultConfigPath() string {
  dirname, _ := os.UserHomeDir()
  return dirname + "/.mdrss/config.json"
}

func ParseArguments(arguments []string) (map[string]*string, error) {
  parsedArguments := make(map[string]*string)
  command, commandErr := getCommand(arguments)
  parsedArguments["config"] = flag.String("config", defaultConfigPath(), "path to config.json")
  parsedArguments["command"] = &command
  flag.Parse()
  return parsedArguments, commandErr
}

