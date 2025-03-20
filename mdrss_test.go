package main

import (
  mdrss "github.com/TimoKats/mdrss/lib"
  "testing"
)

func testConfig() mdrss.Config {
  var config mdrss.Config
  config.Description = "Testing weblog"
  config.InputFolder = "test/"
  config.OutputFile = "rss.xml"
  config.Author = "Testing Test"
  config.Link = "test@testing.com"
  return config
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
  articles := mdrss.ParseMarkdown(config, files)
  got := articles[0].Title
  want := "This article has a title." 
  if got != want {
    t.Errorf("got %v, wanted %v", got, want)
  }
}

func TestConvertLists(t *testing.T) {
  want := "<ul>\n<li>text</li>\n</ul>\n"
  got := mdrss.ConvertMarkdownToXml([]byte("  - text"))
  if got != want {
    t.Errorf("got %v, wanted %v", got, want)
  }
}

func TestConvertLinks(t *testing.T) {
  want := "<p>Click here for the <a href=\"https://timokats.xyz\">link</a></p>\n"
  got := mdrss.ConvertMarkdownToXml([]byte("Click here for the [link](https://timokats.xyz)"))
  if got != want {
    t.Errorf("got %v, wanted %v", got, want)
  }
}

