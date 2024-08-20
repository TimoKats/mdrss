package main

import (
  mdrss "github.com/TimoKats/mdrss/lib"
  "flag"
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
  var config_path string
  flag.StringVar(&config_path, "config", "~/.mdrss/config.json", "path to config.json")
  flag.Parse()

  var cmd string
  switch len(os.Args) {
    case 2:
      cmd = os.Args[1]
    case 4:
      cmd = os.Args[3]
    default:
      mdrss.Error.Println("mdrss <<update, ls, conf >>")
      return
  }

  config, configErr := mdrss.ReadConfig(config_path)
  if configErr != nil {
    mdrss.Error.Println(configErr)
    return
  }
  commandErr := parseCommand(cmd, config)
  if commandErr != nil {
    mdrss.Error.Println(commandErr)
  }
}

