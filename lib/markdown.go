package lib

import (
  "os"
  "bufio"
  "errors"
  "strings"
  "unicode"

  "github.com/gomarkdown/markdown"
  "github.com/gomarkdown/markdown/html"
)

func checkMarkdownTitle(text string) bool {
  if len(text) > 0 {
    return string(text[0]) == string("#")
  }
  return false
}

func getLeadingWhitespace(text string) int {
  for index, letter := range text {
    if unicode.IsLetter(letter) {
      return index
    }
  }
  return 0
}

func ConvertMarkdownToXml(text []byte) string {
  opts := html.RendererOptions{Flags: html.CommonFlags}
  renderer := html.NewRenderer(opts)
  html := markdown.ToHTML(text, nil, renderer)
  return string(html)
}

func GetArticles(config Config) ([]Article, error) {
  RawArticles, fileErr := os.ReadDir(config.InputFolder)
  articles := []Article{}
  if fileErr == nil {
    for _, file := range RawArticles {
      if !file.IsDir() && !strings.HasPrefix(file.Name(), "draft-") && strings.HasSuffix(file.Name(), ".md") {
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
    articleBody := []byte("")
    filePath := config.InputFolder + "/" + articles[index].Filename
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
