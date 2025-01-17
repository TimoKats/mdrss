package lib

import (
  "os"
  "io/fs"
  "bufio"
  "errors"
  "strings"

  "github.com/gomarkdown/markdown"
  "github.com/gomarkdown/markdown/html"
)

func formatGuid(file fs.DirEntry) string {
  filename := strings.Split(file.Name(), ".")[0]
  return strings.ReplaceAll(filename, " ", "-")
}

func validFilename(file fs.DirEntry) bool {
  return !strings.HasPrefix(file.Name(), "draft-") && strings.HasSuffix(file.Name(), ".md")
}

func checkMarkdownTitle(text string) bool {
  if len(text) > 0 {
    return string(text[0]) == string("#")
  }
  return false
}

func ConvertMarkdownToXml(text []byte) string {
  opts := html.RendererOptions{Flags: html.CommonFlags}
  renderer := html.NewRenderer(opts)
  html := markdown.ToHTML(text, nil, renderer)
  return string(html)
}

func GetArticles(feed Feed) ([]Article, error) {
  RawArticles, fileErr := os.ReadDir(feed.InputFolder)
  articles := []Article{}
  if fileErr == nil {
    for _, file := range RawArticles {
      if !file.IsDir() && validFilename(file) {
        var article Article
        fileInfo, _ := file.Info()
        article.Filename = file.Name()
        article.DatePublished = fileInfo.ModTime()
        article.Guid = feed.Link + "/" + formatGuid(file)
        articles = append(articles, article)
      }
    }
    return articles, nil
  }
  return articles, errors.New("Error when getting files from input directory.")
}

func ReadMarkdown(feed Feed, articles []Article) []Article {
  for index := range articles {
    articleBody := []byte("")
    filePath := feed.InputFolder + "/" + articles[index].Filename
    readFile, _ := os.Open(filePath)
    scanner := bufio.NewScanner(readFile)
    for scanner.Scan() {
      if checkMarkdownTitle(scanner.Text()) && len(articles[index].Title) == 0 {
        articles[index].Title = scanner.Text()[2:len(scanner.Text())]
      } else {
        articleBody = append(articleBody, scanner.Text()...)
        articleBody = append(articleBody, "\n"...)
      }
      articles[index].Description = ConvertMarkdownToXml(articleBody)
    }

  }
  return articles
}
