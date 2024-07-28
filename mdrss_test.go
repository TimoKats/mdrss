package main

import (
  mdrss "mdrss/lib"
  "testing"
)

func testConfig() mdrss.Config {
  var config mdrss.Config
  config.Description = "Testing weblog"
  config.InputFolder = "test/"
  config.OutputFile = "index.xml"
  config.Author = "Testing Test"
  config.Link = "test@testing.com"
  return config
}

func TestFileExists(t *testing.T) {
  got := mdrss.FileExists("main.go")
  want := true 
  if got != want {
    t.Errorf("got %v, wanted %v", got, want)
  }
}

func TestGetArticles(t *testing.T) {
  config := testConfig()
  articles, _ := mdrss.GetArticles(config)
  got := []string{}
  want := []string{"another-article.md"}
  for _, article := range articles {
    got = append(got, article.Filename)
  }
  if len(got) != len(want) {
    t.Errorf("got %v, wanted %v", got, want)
  }
}

func TestCreateMarkdown(t *testing.T) {
  config := testConfig()
  files, _ := mdrss.GetArticles(config)
  config.Articles = mdrss.ReadMarkdown(config, files)
  got := config.Articles[0].Title
  want := "This article has a title." 
  if got != want {
    t.Errorf("got %v, wanted %v", got, want)
  }
}

func TestCreateRSS(t *testing.T) {
  config := testConfig()
  files, _ := mdrss.GetArticles(config)
  config.Articles = mdrss.ReadMarkdown(config, files)
  got := len(mdrss.CreateRSS(config))
  want := 481 
  if got != want {
    t.Errorf("got %v, wanted %v", got, want)
  }
}

