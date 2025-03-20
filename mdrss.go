// contains the control logic of mdrss. Parses config/arguments, and calls the required
// function (update, ls) from the mdrss module.

package main

import (
  mdrss "github.com/TimoKats/mdrss/lib"
  "errors"
  "os"
)

func ls(config mdrss.Config) error {
  files, fileErr := mdrss.GetArticles(config)
  if fileErr == nil {
    for _, file := range files {
      mdrss.Info.Println(file)
    }
  }
  return fileErr
}

func update(config mdrss.Config) error {
  var feed mdrss.Feed
  markdownErr := feed.FromConfig(config)
  if markdownErr != nil { return markdownErr }
  rssErr := feed.ToXML()
  if rssErr != nil { return rssErr }
  mdrss.Info.Printf("Content written to %s", feed.Conf.OutputFile)
  return nil
}

func parse(command string, config mdrss.Config) error {
  switch (command) {
    case "ls":
      return ls(config)
    case "update":
      return update(config)
    default:
      return errors.New("Command not found.")
  }
}

func main() {
  // get arguments and config
  arguments, argumentErr := mdrss.ParseArguments(os.Args)
  config, configErr := mdrss.ReadConfig(*arguments["config"])
  mdrss.Info.Printf("Using config location: %v", *arguments["config"])
  if err := errors.Join(argumentErr, configErr); err != nil {
    mdrss.Error.Println(err); return
  }
  // run command
  commandErr := parse(*arguments["command"], config)
  if commandErr != nil {
    mdrss.Error.Println(commandErr)
  }
}

