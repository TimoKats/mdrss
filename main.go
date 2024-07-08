package main

import (
  mdrss "mdrss/lib"
  "os"
)

func parseCommand(command string, config mdrss.Config) error {
  files := mdrss.GetFiles(config)
  articles := mdrss.ReadMarkdown(config, files)
  config.Articles = articles
  rssXml := mdrss.CreateRSS(config)
  mdrss.WriteRSS(rssXml, config)
  return nil
}

func main() {
  if len(os.Args) != 2 { mdrss.Error.Println("mdrss <<update, ls>>"); return }
  configPath := mdrss.GetConfigPath()
  if mdrss.FileExists(configPath) {
    config, _ := mdrss.ReadConfig(configPath)
    parseCommand(os.Args[1], config)
  } else {
    mdrss.Error.Println("No ~/.mdrss/config.json found.")
  }
}

