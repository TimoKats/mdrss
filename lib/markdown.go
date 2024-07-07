package lib

import (
  "io/ioutil"
  "bufio"
  "os"
)

func checkMarkdownTitle(text string) bool {
  if len(text) > 0 {
    return string(text[0]) == string("#")
  }
  return false
}

func GetFiles(config Config) []string {
  RawFiles, _ := ioutil.ReadDir(config.InputFolder)
  files := []string{}
  for _, file := range RawFiles {
    if ! file.IsDir() {
      files = append(files, file.Name())
    }
  }
  return files
}

func ReadMarkdown(config Config, files []string) []Article {
  var articles []Article 
  for _, file := range files {
    var article Article
    filePath := config.InputFolder + "/" + file
    readFile, _ := os.Open(filePath)
    scanner := bufio.NewScanner(readFile)
    for scanner.Scan() {
      if checkMarkdownTitle(scanner.Text()) && len(article.Title) == 0 {
        article.Title = scanner.Text()[1:len(scanner.Text())]
      } else {
        article.Description += scanner.Text()
      }
    }
    articles = append(articles, article)
  }
  return articles
}

