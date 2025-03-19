package main

import (
  mdrss "github.com/TimoKats/mdrss/lib"
  "errors"
  "os"
)

func parseCommand(command string, config mdrss.Feed) error {
  switch (command) {
    case "ls":
      return mdrss.LsCommand(config)
    case "conf":
      return mdrss.ConfCommand(config)
    case "update":
      return mdrss.UpdateCommand(config)
    default:
      return errors.New("Command not found.")
  }
}

func main() {
  arguments, argumentErr := mdrss.ParseArguments(os.Args)
  config, configErr := mdrss.ReadConfig(*arguments["config"])
  mdrss.Info.Printf("Using config location: %v", *arguments["config"])
  if err := errors.Join(argumentErr, configErr); err != nil && *arguments["command"] != "init" {
    mdrss.Error.Println(err); return
  }
  commandErr := parseCommand(*arguments["command"], config)
  if commandErr != nil {
    mdrss.Error.Println(commandErr)
  }
}

