package lib

import (
  "time"
  "os"
)

func formatTimestamp(timestamp time.Time) string {
  return timestamp.Format(time.RFC1123) 
}


func CreateRSS(config Config) string {
  var xmlContent string

  xmlContent = "<?xml version=\"1.0\" encoding=\"UTF-8\" ?>\n"
  xmlContent += "<rss version=\"2.0\">\n"
  xmlContent += "<channel>\n"
  xmlContent += "<title>" + config.Description + "</title>\n"
  xmlContent += "<link>" + config.OutputFile + "</link>\n"
  xmlContent += "<description>" + config.Description + "</description>\n"

  for _, article := range config.Articles {
    if len(article.Title) == 0 {
      continue
    }
    xmlContent += "\t<item>\n"
    xmlContent += "\t\t<title>" + article.Title + "</title>\n"
    xmlContent += "\t\t<link>" + article.Title + "</link>\n"
    xmlContent += "\t\t<pubDate>" + formatTimestamp(article.DatePublished) + "</pubDate>\n"
    xmlContent += "\t\t<description><![CDATA[" + article.Description + "]]></description>\n"
    xmlContent += "\t</item>\n"
    Info.Printf("Added '%s' to RSS feed. ", article.Title)
  }
  
  xmlContent += "</channel>\n</rss>\n"
  return xmlContent 
}

func WriteRSS(rssContent string, config Config) error {
  rssByte := []byte(rssContent)
  err := os.WriteFile(config.OutputFile, rssByte, 0644)
  return err
}

