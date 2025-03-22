// contains the control logic of mdrss. Parses config/arguments, and calls the required
// function (update, ls) from the mdrss module.

package main

import (
  mdrss "github.com/TimoKats/mdrss/lib"
  "errors"
  "os"
)

func ls(config mdrss.Config) error {
  var feed mdrss.Feed
  err := feed.FromConfig(config)
  if err != nil { return err }
  for _, article := range feed.Articles {
    mdrss.Info.Printf("%s/%s", article.Topic, article.Filename)
  }
  return err
}

func update(config mdrss.Config) error {
  var feed mdrss.Feed
  markdownErr := feed.FromConfig(config)
  if markdownErr != nil { return markdownErr }
  rssErr := feed.ToXML()
  if rssErr != nil { return rssErr }
  mdrss.Info.Printf("Content written to %s", config.OutputFile)
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

