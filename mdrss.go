// contains the control logic of mdrss. Parses config/arguments, and calls the required
// function (update, ls) from the mdrss module.

package main

import (
	"errors"
	"os"

	mdrss "github.com/TimoKats/mdrss/pkg"
)

func ls(config mdrss.Config) error {
	var feed mdrss.Feed
	if err := feed.FromConfig(config); err != nil {
		return err
	}
	for _, article := range feed.Articles {
		mdrss.Info.Printf("%s/%s", article.Topic, article.Filename)
	}
	return nil
}

func update(config mdrss.Config) error {
	var feed mdrss.Feed
	if err := feed.FromConfig(config); err != nil {
		return err
	}
	return feed.ToXML()
}

func parse(command string, config mdrss.Config) error {
	switch command {
	case "ls":
		return ls(config)
	case "update":
		return update(config)
	default:
		return errors.New("Command not found.")
	}
}

func main() {
	// get arguments and config
	arguments, argumentErr := mdrss.ParseArguments(os.Args)
	config, configErr := mdrss.ReadConfig(*arguments["config"])
	mdrss.Info.Printf("Using config location: %v", *arguments["config"])
	if err := errors.Join(argumentErr, configErr); err != nil {
		mdrss.Error.Println(err)
		return
	}
	// run command
	commandErr := parse(*arguments["command"], config)
	if commandErr != nil {
		mdrss.Error.Println(commandErr)
	}
}
