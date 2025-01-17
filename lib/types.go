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

type Feed struct {
  Description string
  InputFolder string
  OutputFile string
  Author string
  Link string

  // optional fields
  Articles []Article
}

