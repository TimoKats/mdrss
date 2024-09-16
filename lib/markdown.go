package lib

import (
  "os"
  "bufio"
  "errors"
  "regexp"
  "strconv"
  "strings"
  "unicode"
)

var markdownUlListActive bool
var markdownOlIndex int64
var codeBlockActive = false
var codeBlockAggregate = ""

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

func parseLink(text string, markdownLinks *regexp.Regexp) string {
  if strings.Contains(text, "audio/mpeg") {
    return ConvertEnclosure(text, markdownLinks)
  } else {
    return ConvertLink(text, markdownLinks)
  }
}

func parseCodeblock(text string) string {
  if !codeBlockActive {
    codeBlockActive = true
    codeBlockAggregate = "<sup>" + text[3:] + "</sup><br>"
    return "" + "<pre style=\"word-wrap: break-word;\"><code>"
  } else {
    out := codeBlockAggregate
    codeBlockAggregate, codeBlockActive = "", false
    return out + "</code></pre>"
  }
  if codeBlockAggregate != "" {
    codeBlockAggregate += "<br>"
  }
  codeBlockAggregate += text
  return ""
}

func parseOrderedLists(text string, markdownOrderedLists *regexp.Regexp) string {
  entryIndex, entryErr := strconv.ParseInt(markdownOrderedLists.FindStringSubmatch(text)[2], 10, 64)
  entryText := markdownOrderedLists.FindStringSubmatch(text)[4]
  if entryErr != nil {
    return "<p>" + text + "</p>"
  }
  return ConvertOrderedlList(entryText, entryIndex)
}

func ConvertMarkdownToRSS(text string) string {
  markdownLinks := regexp.MustCompile(`\[(.*)\]\((.*)\)`)
  markdownUnorderedLists := regexp.MustCompile(`^(\s*)(-|\*|\+)[\s](.*)`)
  markdownOrderedLists := regexp.MustCompile(`^(\s*)(-?\d+)(\.\s+)(.*)$`)
  fencedCodeBlock := regexp.MustCompile("^\x60\x60\x60")
  text = ConvertTextEnrichment(text)
  switch {
    // links 
    case markdownLinks.MatchString(text):
      return parseLink(text, markdownLinks)
    // lists
    case markdownUnorderedLists.MatchString(text):
      return ConvertUnorderedlList(text)
    case markdownUlListActive:
      markdownUlListActive = false
      return "</ul><p>" + text + "</p>"
    case markdownOrderedLists.MatchString(text):
      return parseOrderedLists(text, markdownOrderedLists)
    case markdownOlIndex != 0:
      markdownOlIndex = 0
      return "</ol>" + ConvertMarkdownToRSS(text)
    //code blocks
    case fencedCodeBlock.MatchString(text) || codeBlockActive:
      return parseCodeblock(text)
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
