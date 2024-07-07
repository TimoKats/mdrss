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

// type Channel struct {
//   Name string
//   Author string
//   Articles []Article
// 
//   // optional fields
//   CopyrightNotice string
// }

type Config struct {
  InputFolder string
  OutputFile string 
  Author string

  // optional fields
  Articles []Article
}

