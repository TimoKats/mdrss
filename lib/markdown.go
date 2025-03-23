// Module responsible for reading markdown files and converting it to valid HTML.
// FromConfig is called by the main control flow. It adds the config first and then the
// articles using a (semi) recursive function. Remaining functions are helpers for conv-
// erting markdown to HMTL, etc.

package lib

import (
  "os"
  "io/fs"
  "bufio"
  "strings"
  "path/filepath"

  "github.com/gomarkdown/markdown"
  "github.com/gomarkdown/markdown/html"
)

func getLastFolder(path string) string {
  cleanedPath := filepath.Clean(path)
  parts := strings.Split(cleanedPath, string(filepath.Separator))
  if len(parts) == 0 {
    Error.Println("No folder found."); return ""
  }
  return parts[len(parts)-1]
}

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

func convertMarkdownToXml(text []byte) string {
  opts := html.RendererOptions{Flags: html.CommonFlags}
  renderer := html.NewRenderer(opts)
  html := markdown.ToHTML(text, nil, renderer)
  return string(html)
}

func parseMarkdown(article Article, config Config) Article {
  articleBody := []byte("")
  filePath := config.InputFolder + "/" + article.Topic + "/" + article.Filename
  readFile, _ := os.Open(filePath)
  scanner := bufio.NewScanner(readFile)
  for scanner.Scan() {
    if checkMarkdownTitle(scanner.Text()) && len(article.Title) == 0 {
      article.Title = scanner.Text()[2:len(scanner.Text())]
    } else {
      articleBody = append(articleBody, scanner.Text()...)
      articleBody = append(articleBody, "\n"...)
    }
    article.Description = convertMarkdownToXml(articleBody)
  }
  return article
}

func newArticle(file fs.DirEntry, config Config) Article {
  fileInfo, _ := file.Info()
  article := Article{
    Filename: file.Name(),
    DatePublished: fileInfo.ModTime(),
    Guid: config.Link + "/" + formatGuid(file),
  }
  if len(config.topicInputFolder) > 0 {
    article.Topic = getLastFolder(config.topicInputFolder)
  }
  return parseMarkdown(article, config)
}

func (feed *Feed) getArticles() error {
  rawArticles, err := os.ReadDir(feed.config.InputFolder)
  if len(feed.config.topicInputFolder) > 0 {
    rawArticles, err = os.ReadDir(feed.config.topicInputFolder)
  }
  if err != nil { return err }
  for _, file := range rawArticles {
    if file.IsDir() {
      feed.config.topicInputFolder = feed.config.InputFolder + "/" + file.Name()
      feed.getArticles()
    } else if !file.IsDir() && validFilename(file) {
      feed.Articles = append(feed.Articles, newArticle(file, feed.config))
    }
  }
  return nil
}

func (feed *Feed) FromConfig(config Config) error {
  feed.config = config
  return feed.getArticles()
}

