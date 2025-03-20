// Module responsible for writing RSS compliant XML based on a feed objcet.

package lib

import (
  "time"
  "os"
)

func addItem(xmlContent string, config Config, article Article) string {
  timestamp := article.DatePublished.Format(time.RFC822Z)
  xmlContent += "\t<item>\n"
  xmlContent += "\t\t<title>" + article.Title + "</title>\n"
  xmlContent += "\t\t<link>" + config.Link + "</link>\n"
  xmlContent += "\t\t<pubDate>" + timestamp + "</pubDate>\n"
  xmlContent += "\t\t<guid>" + article.Guid + "</guid>\n"
  xmlContent += "\t\t<description><![CDATA[" + article.Description + "]]></description>\n"
  xmlContent += "\t</item>\n"
  return xmlContent
}

func addHeader(config Config) string {
  timestamp := time.Now().Format(time.RFC822Z)
  xmlContent := "<?xml version=\"1.0\" encoding=\"UTF-8\" ?>\n"
  xmlContent += "<rss version=\"2.0\">\n"
  xmlContent += "<channel>\n"
  xmlContent += "<title>" + config.Description + "</title>\n"
  xmlContent += "<author>" + config.Author + "</author>\n"
  xmlContent += "<link>" + config.Link + "</link>\n"
  xmlContent += "<lastBuildDate>" + timestamp + "</lastBuildDate>\n"
  xmlContent += "<description>" + config.Description + "</description>\n"
  return xmlContent
}

func (feed *Feed) ToXML() error {
  xmlContent := addHeader(feed.Conf)
  for _, article := range feed.Articles {
    if len(article.Title) != 0 {
      xmlContent = addItem(xmlContent, feed.Conf, article)
      Info.Printf("Added '%s' to RSS config. ", article.Title)
    } else {
      Warn.Printf("%s doesn't have a valid markdown title.", article.Filename)
    }
  }
  xmlContent += "</channel>\n</rss>\n"
  rssByte := []byte(xmlContent)
  return os.WriteFile(feed.Conf.OutputFile, rssByte, 0644)
}

