package main

import (
  "os"
  mdrss "mdrss/lib"
)

func main() {
  if len(os.Args) < 2 {
    mdrss.Error.Println("No path to config.json given.")
    return
  }
  configFile := os.Args[1]
  if mdrss.FileExists(configFile) {
    config, _ := mdrss.ReadConfig(configFile)
    files := mdrss.GetFiles(config)
    articles := mdrss.ReadMarkdown(config, files)
    config.Articles = articles
    rssXml := mdrss.CreateRSS(config)
    mdrss.WriteRSS(rssXml, config)
  } else {
    mdrss.Error.Println("No config.json found.")
  }
}

