package lib

import (
  "strings"
  "errors"
  "bufio"
  "os"
)

func checkMarkdownTitle(text string) bool {
  if len(text) > 0 {
    return string(text[0]) == string("#")
  }
  return false
}

func GetArticles(config Config) ([]Article, error) {
  RawArticles, fileErr := os.ReadDir(config.InputFolder)
  articles := []Article{}
  if fileErr == nil {
    for _, file := range RawArticles {
      if !file.IsDir() && !strings.HasPrefix(file.Name(), "draft-") {
        var article Article
        fileInfo, _ := file.Info()
        article.Filename = file.Name() 
        article.DatePublished = fileInfo.ModTime()
        articles = append(articles, article)
      }
    }
    return articles, nil
  }
  return articles, errors.New("Error when getting files from input directory.")
}

func ReadMarkdown(config Config, articles []Article) []Article {
  for index := range articles {
    filePath := config.InputFolder + "/" + articles[index].Filename 
    readFile, _ := os.Open(filePath)
    scanner := bufio.NewScanner(readFile)
    for scanner.Scan() {
      if checkMarkdownTitle(scanner.Text()) && len(articles[index].Title) == 0 {
        articles[index].Title = scanner.Text()[2:len(scanner.Text())]
      } else if len(scanner.Text()) > 0 {
        articles[index].Description += "<p>" + scanner.Text() + "</p>"
      }
    }
  }
  return articles
}

