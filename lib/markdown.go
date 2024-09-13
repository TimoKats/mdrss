package lib

import (
  "bufio"
  "errors"
  "os"
  "regexp"
  "strconv"
  "strings"
  "unicode"
)

var markdownUlListActive bool
var markdownOlIndex int64

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
  url := markdownLinks.ReplaceAllString(text, "$2")
  size, fileSizeErr := FileSizeUrl(url)
  if fileSizeErr != nil {
    Error.Println(fileSizeErr)
    return ""
  }
  return "<enclosure " + markdownLinks.ReplaceAllString(text, "url='$2' type='$1' length='") + size + "' />"
}

func convertMarkdownUlList(text string) string {
  if !markdownUlListActive {
    markdownUlListActive = true
    return "<ul><li>" + text[getLeadingWhitespace(text):] + "</li>"
  }
  return "<li>" + text[getLeadingWhitespace(text):] + "</li>"
}

func convertMarkdownOlList(text string, index int64) string {
  if markdownOlIndex == 0 {
    markdownOlIndex = 1
    return "<ol><li>" + text[getLeadingWhitespace(text):] + "</li>"
  }
  return "<li>" + text[getLeadingWhitespace(text):] + "</li>"
}

func ConvertMarkdownToRSS(text string) string {
  markdownLinks := regexp.MustCompile(`\[(.*)\]\((.*)\)`)
  markdownUnorderedLists := regexp.MustCompile(`^(\s*)(-|\*|\+)[\s](.*)`)
  markdownOrderedLists := regexp.MustCompile(`^(\s*)(-?\d+)(\.\s+)(.*)$`)
  if markdownLinks.Match([]byte(text)) && strings.Contains(text, "audio/mpeg") {
    return convertMarkdownEnclosure(text, markdownLinks)
  }
  if markdownLinks.Match([]byte(text)) {
    return convertMarkdownLink(text, markdownLinks)
  }
  if markdownUnorderedLists.Match([]byte(text)) {
    return convertMarkdownUlList(text)
  } else if markdownUlListActive {
    markdownUlListActive = false
    return "</ul><p>" + text + "</p>"
  }
  if markdownOrderedLists.Match([]byte(text)) {
    entryIndex, entryErr := strconv.ParseInt(markdownOrderedLists.FindStringSubmatch(text)[2], 10, 64)
    entryText := markdownOrderedLists.FindStringSubmatch(text)[4]
    if entryErr != nil {
      return "<p>" + text + "</p>"
    }
    return convertMarkdownOlList(entryText, entryIndex)
  } else if markdownOlIndex != 0 {
    markdownOlIndex = 0
    return "</ol><p>" + text + "</p>"
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
    markdownUlListActive = false
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
