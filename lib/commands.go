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

func LsCommand(config Config) error {
  files, fileErr := GetArticles(config)
  if fileErr == nil {
    for _, file := range files {
      Info.Println(file)
    }
  }
  return fileErr
}

func UpdateCommand(config Config) error {
  files, fileErr := GetArticles(config)
  if fileErr == nil {
    config.Articles = ReadMarkdown(config, files)
    rssXml := CreateRSS(config)
    rssErr := WriteRSS(rssXml, config)
    if rssErr != nil {
      return rssErr
    }
    Info.Printf("Content written to %s", config.OutputFile)
  }
  return fileErr
}

func ConfCommand(config Config) error {
	configValues := reflect.ValueOf(config)
  typeOfS := configValues.Type()
  for i := 0; i < configValues.NumField(); i++ {
    if len(typeOfS.Field(i).Name) < 8 {
      Info.Printf("%s\t\t%v\n", typeOfS.Field(i).Name, configValues.Field(i).Interface())
    } else {
      Info.Printf("%s\t%v\n", typeOfS.Field(i).Name, configValues.Field(i).Interface())
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
  var configPath string
  var newConfig Config
  promptUsr("Feed description", &newConfig.Description, "An RSS feed")
  promptUsr("Feed link", &newConfig.Link, "https://localhost:5500/index.xml")
  promptUsr("Feed author", &newConfig.Author, "anonymous")
  promptUsr("Config path", &configPath, DefaultConfigPath())
  promptUsr("Input folder path (where you have md files)", &newConfig.InputFolder, "")
  promptUsr("Output file path", &newConfig.OutputFile, "index.xml")

  // Only update file after ALL init fields are inputted
  dirErr := os.MkdirAll(filepath.Dir(configPath), 0755)
  configFile, fileErr := os.Create(configPath)
  out, jsonErr := json.MarshalIndent(newConfig, "", "  ")
  if err := errors.Join(dirErr, fileErr, jsonErr); err != nil {
    return err
  }
  _, writeErr := configFile.Write(out)
  if writeErr == nil {
    Info.Println("Initialized! Run \"mdrss update\" to update your feed!")
  }
  return writeErr
}
