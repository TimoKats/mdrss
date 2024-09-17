package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"strings"

	mdrss "github.com/TimoKats/mdrss/lib"
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

func initCommand() error {
	configPath := mdrss.DefaultConfigPath()
	var newConfig mdrss.Config
	promptUsr("Feed title", &newConfig.Title, "Untitled")
	promptUsr("Feed link", &newConfig.Link, "https://localhost:5500/index.xml")
	promptUsr("Feed description", &newConfig.Description, "An RSS feed")
	promptUsr("Feed author", &newConfig.Channel.Author, "anonymous")
	defaultDir, _ := os.Getwd()
	promptUsr("Input folder path", &newConfig.InputFolder, defaultDir+"/test/")
	promptUsr("Output file path", &newConfig.OutputFile, defaultDir+"/index.xml")

	// Only update file after ALL init fields are inputted
	dirCreateErr := os.MkdirAll(filepath.Dir(configPath), 0755)
	configFile, fileCreateErr := os.Create(configPath)
	if fileCreateErr != nil {
		if dirCreateErr != nil {
			mdrss.Error.Println("unable to create configuration directory")
		} else {
			mdrss.Error.Println("unable to create configuration file")
		}
		return nil
	}

	out, jsonErr := json.MarshalIndent(newConfig, "", "    ")
	if jsonErr != nil {
		mdrss.Error.Println("Critical failure in initialization of config")
	}
	_, writeErr := configFile.Write(out)
	fmt.Println("\033[32m MDRSS Initialized! Run \"mdrss update\" to update your feed!")
	return writeErr
}

func parseCommand(command string, config mdrss.Config) error {
	switch command {
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
	if *arguments["command"] == "init" {
		initErr := initCommand()
		if initErr != nil {
			mdrss.Error.Println(initErr)
		}
		return
	}
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
