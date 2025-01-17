package main

import (
  mdrss "github.com/TimoKats/mdrss/lib"
  "os/exec"
  "testing"
)

func testFeed() mdrss.Feed {
  var feed mdrss.Feed
  feed.Description = "Testing weblog"
  feed.InputFolder = "test/"
  feed.OutputFile = "rss.xml"
  feed.Author = "Testing Test"
  feed.Link = "test@testing.com"
  return feed
}

func TestFileExists(t *testing.T) {
  got := mdrss.FileExists("mdrss.go")
  want := true
  if got != want {
    t.Errorf("got %v, wanted %v", got, want)
  }
}

func TestGetArticles(t *testing.T) {
  feed := testFeed()
  articles, _ := mdrss.GetArticles(feed)
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
  feed := testFeed()
  files, _ := mdrss.GetArticles(feed)
  feed.Articles = mdrss.ReadMarkdown(feed, files)
  got := feed.Articles[0].Title
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

func TestBasics(t *testing.T) {
  feed := testFeed()
  articles, _ := mdrss.GetArticles(feed)
  feed.Articles = mdrss.ReadMarkdown(feed, articles)
  rssXml := mdrss.CreateRSS(feed)
  rssErr := mdrss.WriteRSS(rssXml, feed)
  got := mdrss.FileExists("rss.xml")
  want := true
  if got != want || rssErr != nil {
    t.Errorf("got %v, wanted %v", got, want)
  } else {
    cmd := exec.Command("rm", "rss.xml")
    runErr := cmd.Run()
    if runErr != nil {
      mdrss.Error.Println(runErr)
    }
  }
}

