package lib

import (
  "regexp"
  "time"
)

type Markdown struct {
  Content []Line
}

type Line struct {
  Link *regexp.Regexp
  UnorderedList *regexp.Regexp
  OrderedList *regexp.Regexp
  CodeBlock *regexp.Regexp

  // optional fields
  CodeBlockOpen bool
  UnorderedListActive bool
  OrderedListActive bool
}

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

