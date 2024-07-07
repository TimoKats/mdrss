package main

import (
  mdrss "mdrss/lib"
)

func main() {
  if mdrss.FileExists("config.json") {
    config, _ := mdrss.ReadConfig("config.json")
    files := mdrss.GetFiles(config)
    articles := mdrss.ReadMarkdown(config, files)
    config.Articles = articles
    rssXml := mdrss.CreateRSS(config)
    mdrss.WriteRSS(rssXml, config)
  } else {
    mdrss.Error.Println("No config.json found.")
  }
}

