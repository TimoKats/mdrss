package main

import (
  mdrss "mdrss/lib"
  "errors"
  "os"
)

func lsCommand(config mdrss.Config) error {
  files, fileErr := mdrss.GetFiles(config)
  if fileErr == nil {
    for _, file := range files {
      mdrss.Info.Println(file)
    }
  }
  return fileErr 
}

func updateCommand(config mdrss.Config) error {
  files, fileErr := mdrss.GetFiles(config)
  if fileErr == nil {
    config.Articles = mdrss.ReadMarkdown(config, files)
    rssXml := mdrss.CreateRSS(config)
    mdrss.WriteRSS(rssXml, config)
    mdrss.Info.Printf("Content written to %s", config.OutputFile)
  }
  return fileErr
}


func parseCommand(command string, config mdrss.Config) error {
  switch (command) {
    case "ls":
      return lsCommand(config)
    case "update":
      return updateCommand(config)
    default:
      return errors.New("Command not found.")
  }
  return nil
}

func main() {
  if len(os.Args) != 2 { mdrss.Error.Println("mdrss <<update, ls>>"); return }
  configPath := mdrss.GetConfigPath()
  config, configErr := mdrss.ReadConfig(configPath)
  if configErr == nil {
    parseCommand(os.Args[1], config)
  } else {
    mdrss.Error.Println(configErr)
  }
}

