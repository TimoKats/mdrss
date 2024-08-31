package lib

import (
  "unicode"
  "strings"
  "regexp"
  "errors"
  "bufio"
  "os"
)

var markdownListActive bool

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

func convertMarkdownLink(text string, markdownLinks *regexp.Regexp) string {
  return "<p>" + markdownLinks.ReplaceAllString(text, "<a href='$2'>$1</a>") + "</p>"
}

func convertMarkdownEnclosure(text string, markdownLinks *regexp.Regexp) string {
  return "<enclosure " + markdownLinks.ReplaceAllString(text, "url='$2' type='$1'") + " />"
}

func convertMarkdownList(text string) string {
  if !markdownListActive {
    markdownListActive = true
    return "<ul><li>" + text[getLeadingWhitespace(text):] + "</li>"
  }
  return "<li>" + text[getLeadingWhitespace(text):] + "</li>"
}

func ConvertMarkdownToRSS(text string) string {
  markdownLinks := regexp.MustCompile(`\[(.*)\]\((.*)\)`)
  markdownLists := regexp.MustCompile(`^(\s*)(-|\*|\+)(.*)`)
  if markdownLinks.Match([]byte(text)) && strings.Contains(text, "audio/mpeg") {
    return convertMarkdownEnclosure(text, markdownLinks)
  }
  if markdownLinks.Match([]byte(text)) {
    return convertMarkdownLink(text, markdownLinks)
  }
  if markdownLists.Match([]byte(text)) {
    return convertMarkdownList(text)
  } else if markdownListActive {
    markdownListActive = false
    return "</ul><p>" + text + "</p>"
  }
  return "<p>" + text + "</p>"
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
    markdownListActive = false
    filePath := config.InputFolder + "/" + articles[index].Filename 
    readFile, _ := os.Open(filePath)
    scanner := bufio.NewScanner(readFile)
    for scanner.Scan() {
      if checkMarkdownTitle(scanner.Text()) && len(articles[index].Title) == 0 {
        articles[index].Title = scanner.Text()[2:len(scanner.Text())]
      } else if len(scanner.Text()) > 0 {
        articles[index].Description += ConvertMarkdownToRSS(scanner.Text())
      }
    }
  }
  return articles
}

