package main

import (
  mdrss "github.com/TimoKats/mdrss/lib"
  "errors"
  "os"
)

func lsCommand(config mdrss.Config) error {
  files, fileErr := mdrss.GetArticles(config)
  if fileErr == nil {
    for _, file := range files {
      mdrss.Info.Println(file)
    }
  }
  return fileErr
}

func updateCommand(config mdrss.Config) error {
  files, fileErr := mdrss.GetArticles(config)
  if fileErr == nil {
    config.Articles = mdrss.ReadMarkdown(config, files)
    rssXml := mdrss.CreateRSS(config)
    rssErr := mdrss.WriteRSS(rssXml, config)
    if rssErr != nil {
      return rssErr
    }
    mdrss.Info.Printf("Content written to %s", config.OutputFile)
  }
  return fileErr
}

func confCommand(config mdrss.Config) error {
  mdrss.Info.Println("MDRSS v0.0.3 by Timo Kats")
  mdrss.Info.Println(config.InputFolder)
  return nil
}

func parseCommand(command string, config mdrss.Config) error {
  switch (command) {
    case "ls":
      return lsCommand(config)
    case "conf":
      return confCommand(config)
    case "update":
      return updateCommand(config)
    default:
      return errors.New("Command not found.")
  }
}

func main() {
  if len(os.Args) != 2 {
    mdrss.Error.Println("mdrss <<update, ls, conf >>")
    return
  }
  config, configErr := mdrss.ReadConfig()
  if configErr != nil {
    mdrss.Error.Println(configErr)
    return
  }
  commandErr := parseCommand(os.Args[1], config)
  if commandErr != nil {
    mdrss.Error.Println(commandErr)
  }
}

