package lib

import (
  "time"
)

type Article struct {
  Id int
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

  // optional fields
  Articles []Article
}

