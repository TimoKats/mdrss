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

