// Types/Interfaces for creating feeds. A feed consists of a config and list of articles.
// The interface has two functions, from markdown, to XML (which makes sense...).

package lib

import (
  "time"
)

type Article struct {
  Id int
  Guid string
  Title string
  Filename string
  Description string
  DatePublished time.Time
}

type Config struct {
  Description string
  InputFolder string
  OutputFile string
  Author string
  Link string
}

type Feed struct {
  Conf Config
  Articles []Article
}

type Feeder interface {
  FromConfig(Config) error
  ToXML() error
}

