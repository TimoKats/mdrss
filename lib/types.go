package lib

import (
  "regexp"
  "time"
)

type Article struct {
  Id            int
  Title         string
  Author        string
  Filename      string
  Description   string
  DatePublished time.Time
}

type Config struct {
  InputFolder string
  OutputFile  string
  // Required for valid RSS
  Title       string
  Link        string
  Description string

  // optional fields
  Articles []Article
  Channel  ChannelAttributes
}

type ChannelAttributes struct {
  // https://www.rssboard.org/rss-specification
  Author         string
  Language       string
  Copyright      string
  ManagingEditor string
  WebMaster      string
  Category       string
  Generator      string
  Docs           string
  Cloud          string
  Ttl            string
  Image          string
  Rating         string
  SkipHours      string
  SkipDays       string
}
