package lib

import (
  "os"
)

func CreateRSS(config Config) string {
  var xmlContent string

  xmlContent = "<?xml version=\"1.0\" encoding=\"UTF-8\" ?>\n"
  xmlContent += "<rss version=\"2.0\">\n"
  xmlContent += "<channel>\n"
  xmlContent += "<title>" + config.Author + "</title>\n"
  xmlContent += "<link>" + config.OutputFile + "</link>\n"
  xmlContent += "<description>" + config.Description + "</description>\n"

  for _, article := range config.Articles {
    xmlContent += "\t<item>\n"
    xmlContent += "\t\t<title>" + article.Title + "</title>\n"
    xmlContent += "\t\t<link>" + article.Title + "</link>\n"
    xmlContent += "\t\t<description>" + article.Description + "</description>\n"
    xmlContent += "\t</item>\n"
  }
  
  xmlContent += "</channel>\n</rss>"
  return xmlContent 
}

func WriteRSS(rssContent string, config Config) error {
  rssByte := []byte(rssContent)
  err := os.WriteFile(config.OutputFile, rssByte, 0644)
  return err
}

