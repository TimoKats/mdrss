package lib

import (
  "time"
)

type Article struct {
  Id int
  Title string
  Description string
  DatePublished time.Time

  // optional fields
  Language string
}

type Config struct {
  Description string
  InputFolder string
  OutputFile string
  Author string

  // optional fields
  Articles []Article
}

