package lib

import (
  "os"
  "bufio"
  "errors"
  "regexp"
  "strings"
  "unicode"
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

func ConvertMarkdownToRSS(text string) string {
  markdownLinks := regexp.MustCompile(`\[(.*)\]\((.*)\)`)
  markdownUnorderedLists := regexp.MustCompile(`^(\s*)(-|\*|\+)[\s](.*)`)
  markdownOrderedLists := regexp.MustCompile(`^(\s*)(-?\d+)(\.\s+)(.*)$`)
  fencedCodeBlock := regexp.MustCompile("^\x60\x60\x60")
  text = ConvertTextEnrichment(text)
  switch {
    // links 
    case markdownLinks.MatchString(text) && !codeBlockActive:
      return ConvertLink(text, markdownLinks)
    // lists
    case (markdownUnorderedLists.MatchString(text) || markdownUlListActive) && !codeBlockActive:
      return ConvertUnorderedlList(text, markdownUnorderedLists)
    case (markdownOrderedLists.MatchString(text) || markdownOlIndex != 0) && !codeBlockActive:
      return ConvertOrderedLists(text, markdownOrderedLists)
    //code blocks
    case fencedCodeBlock.MatchString(text) || codeBlockActive:
      return ConvertCodeblock(text, fencedCodeBlock)
  default:
    return "<p>" + text + "</p>"
  }
}

func inlineRewrap(text string, pattern *regexp.Regexp, prefix string, postfix string) string {
  if pattern.Match([]byte(text)) {
    out := pattern.ReplaceAllFunc([]byte(text), func(b []byte) []byte {
      return []byte(prefix + pattern.FindStringSubmatch(string(b))[2] + postfix)
    })
    return string(out)
  } else {
    return text
  }
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
      } else if len(scanner.Text()) > 0 || codeBlockActive {
        articles[index].Description += ConvertMarkdownToRSS(scanner.Text())
      }
    }
  }
  return articles
}
