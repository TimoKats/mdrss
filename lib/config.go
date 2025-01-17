package lib

import (
	"encoding/json"
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

func DefaultConfigPath() string {
  dirname, _ := os.UserHomeDir()
  return dirname + "/.mdrss/config.json"
}

func ParseArguments(arguments []string) (map[string]*string, error) {
  parsedArguments := make(map[string]*string)
  command, commandErr := getCommand(arguments)
  parsedArguments["config"] = flag.String("config", DefaultConfigPath(), "path to feed.json")
  parsedArguments["command"] = &command
  flag.Parse()
  return parsedArguments, commandErr
}

func FileExists(filename string) bool {
  if _, err := os.Stat(filename); err != nil {
    return false
  }
  return true
}

func ReadConfig(feedPath string) (Feed, error) {
  var feed Feed
  if FileExists(feedPath) {
    feedContent, _ := os.ReadFile(feedPath)
    jsonErr := json.Unmarshal(feedContent, &feed)
    if jsonErr != nil {
      return feed, errors.New("Error when reading feed file.")
    }
    return feed, nil
  }
  return feed, errors.New("Feed file not found. Please add it at ~/.mdrss/feed.json")
}

