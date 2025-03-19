package lib

import (
  "reflect"
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

