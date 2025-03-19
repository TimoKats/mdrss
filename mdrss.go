package main

import (
  mdrss "github.com/TimoKats/mdrss/lib"
  "errors"
  "os"
)

func parseCommand(command string, feed mdrss.Feed) error {
  switch (command) {
    case "ls":
      return mdrss.LsCommand(feed)
    case "conf":
      return mdrss.ConfCommand(feed)
    case "update":
      return mdrss.UpdateCommand(feed)
    default:
      return errors.New("Command not found.")
  }
}

func main() {
  arguments, argumentErr := mdrss.ParseArguments(os.Args)
  feed, feedErr := mdrss.ReadConfig(*arguments["config"])
  mdrss.Info.Printf("Using config location: %v", *arguments["config"])
  if err := errors.Join(argumentErr, feedErr); err != nil && *arguments["command"] != "init" {
    mdrss.Error.Println(err); return
  }
  commandErr := parseCommand(*arguments["command"], feed)
  if commandErr != nil {
    mdrss.Error.Println(commandErr)
  }
}

