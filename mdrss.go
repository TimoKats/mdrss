package main

import (
  mdrss "github.com/TimoKats/mdrss/lib"
  "reflect"
  "errors"
  "os"
)

func lsCommand(config mdrss.Config) error {
  files, fileErr := mdrss.GetArticles(config)
  if fileErr == nil {
    for _, file := range files {
      mdrss.Info.Println(file)
    }
  }
  return fileErr
}

func updateCommand(config mdrss.Config) error {
  files, fileErr := mdrss.GetArticles(config)
  if fileErr == nil {
    config.Articles = mdrss.ReadMarkdown(config, files)
    rssXml := mdrss.CreateRSS(config)
    rssErr := mdrss.WriteRSS(rssXml, config)
    if rssErr != nil {
      return rssErr
    }
    mdrss.Info.Printf("Content written to %s", config.OutputFile)
  }
  return fileErr
}

func confCommand(config mdrss.Config) error {
	configValues := reflect.ValueOf(config)
  typeOfS := configValues.Type()
  for i := 0; i < configValues.NumField(); i++ {
    if len(typeOfS.Field(i).Name) < 8 {
      mdrss.Info.Printf("%s\t\t%v\n", typeOfS.Field(i).Name, configValues.Field(i).Interface())
    } else {
      mdrss.Info.Printf("%s\t%v\n", typeOfS.Field(i).Name, configValues.Field(i).Interface())
    }
  }
  return nil
}

func parseCommand(command string, config mdrss.Config) error {
  switch (command) {
    case "ls":
      return lsCommand(config)
    case "conf":
      return confCommand(config)
    case "update":
      return updateCommand(config)
    default:
      return errors.New("Command not found.")
  }
}

func main() {
  arguments, argumentErr := mdrss.ParseArguments(os.Args)
  config, configErr := mdrss.ReadConfig(*arguments["config"])
  mdrss.Info.Printf("Using config location: %v", *arguments["config"])
  if err := errors.Join(argumentErr, configErr); err != nil {
     mdrss.Error.Println(err)
     return
   }
  commandErr := parseCommand(*arguments["command"], config)
  if commandErr != nil {
    mdrss.Error.Println(commandErr)
  }
}

