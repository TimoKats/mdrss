// Types/Interfaces for creating feeds. A feed consists of a config and list of articles.
// The interface has public two functions, from config, to XML (which makes sense...).

package lib

import (
	"time"
)

type Article struct {
	Id            int
	Guid          string
	Title         string
	Filename      string
	Description   string
	DatePublished time.Time

	// optional
	Topic string
}

type Config struct {
	Description string
	InputFolder string
	Author      string
	Link        string

	// optional (one is required)
	OutputFolder string
	OutputFile   string
}

type Feed struct {
	Articles []Article

	// private
	config Config
	topics []string
}

