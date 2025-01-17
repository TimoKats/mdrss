package lib

import (
  "bufio"
  "reflect"
  "encoding/json"
  "errors"
  "os"
  "fmt"
  "path/filepath"
  "strings"

)

func LsCommand(feed Feed) error {
  files, fileErr := GetArticles(feed)
  if fileErr == nil {
    for _, file := range files {
      Info.Println(file)
    }
  }
  return fileErr
}

func UpdateCommand(feed Feed) error {
  files, fileErr := GetArticles(feed)
  if fileErr == nil {
    feed.Articles = ReadMarkdown(feed, files)
    rssXml := CreateRSS(feed)
    rssErr := WriteRSS(rssXml, feed)
    if rssErr != nil { return rssErr }
    Info.Printf("Content written to %s", feed.OutputFile)
  }
  return fileErr
}

func ConfCommand(feed Feed) error {
	feedValues := reflect.ValueOf(feed)
  typeOfS := feedValues.Type()
  for i := 0; i < feedValues.NumField(); i++ {
    if len(typeOfS.Field(i).Name) < 8 {
      Info.Printf("%s\t\t%v\n", typeOfS.Field(i).Name, feedValues.Field(i).Interface())
    } else {
      Info.Printf("%s\t%v\n", typeOfS.Field(i).Name, feedValues.Field(i).Interface())
    }
  }
  return nil
}

func promptUsr(prompt string, attrubute *string, fallback string) {
  fmt.Print("\033[33m?\033[0m " + prompt + " \033[37;2m(" + fallback + ") > \033[0m")
  var input string
  in := bufio.NewReader(os.Stdin)
  input, err := in.ReadString('\n')
  input = strings.TrimSpace(input)
  if input == "" || err != nil {
    *attrubute = fallback
    return
  }
  *attrubute = input
}

func InitCommand() error {
  var feedPath string
  var newFeed Feed
  promptUsr("Feed description", &newFeed.Description, "An RSS feed")
  promptUsr("Feed link", &newFeed.Link, "https://localhost:5500/index.xml")
  promptUsr("Feed author", &newFeed.Author, "anonymous")
  promptUsr("Feed path", &feedPath, DefaultConfigPath())
  promptUsr("Input folder path (where you have md files)", &newFeed.InputFolder, "")
  promptUsr("Output file path", &newFeed.OutputFile, "index.xml")

  // Only update file after ALL init fields are inputted
  dirErr := os.MkdirAll(filepath.Dir(feedPath), 0755)
  feedFile, fileErr := os.Create(feedPath)
  out, jsonErr := json.MarshalIndent(newFeed, "", "  ")
  if err := errors.Join(dirErr, fileErr, jsonErr); err != nil {
    return err
  }
  _, writeErr := feedFile.Write(out)
  if writeErr == nil {
    Info.Println("Initialized! Run \"mdrss update\" to update your feed!")
  }
  return writeErr
}
