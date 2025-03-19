package lib

import (
  "time"
  "os"
)

func addItem(xmlContent string, feed Feed, article Article) string {
  timestamp := article.DatePublished.Format(time.RFC822Z)
  xmlContent += "\t<item>\n"
  xmlContent += "\t\t<title>" + article.Title + "</title>\n"
  xmlContent += "\t\t<link>" + feed.Link + "</link>\n"
  xmlContent += "\t\t<pubDate>" + timestamp + "</pubDate>\n"
  xmlContent += "\t\t<guid>" + article.Guid + "</guid>\n"
  xmlContent += "\t\t<description><![CDATA[" + article.Description + "]]></description>\n"
  xmlContent += "\t</item>\n"
  return xmlContent
}

func addHeader(feed Feed) string {
  timestamp := time.Now().Format(time.RFC822Z)
  xmlContent := "<?xml version=\"1.0\" encoding=\"UTF-8\" ?>\n"
  xmlContent += "<rss version=\"2.0\">\n"
  xmlContent += "<channel>\n"
  xmlContent += "<title>" + feed.Description + "</title>\n"
  xmlContent += "<author>" + feed.Author + "</author>\n"
  xmlContent += "<link>" + feed.Link + "</link>\n"
  xmlContent += "<lastBuildDate>" + timestamp + "</lastBuildDate>\n"
  xmlContent += "<description>" + feed.Description + "</description>\n"
  return xmlContent
}

func CreateRSS(feed Feed) string {
  xmlContent := addHeader(feed)
  for _, article := range feed.Articles {
    if len(article.Title) != 0 {
      xmlContent = addItem(xmlContent, feed, article)
      Info.Printf("Added '%s' to RSS feed. ", article.Title)
    } else {
      Warn.Printf("%s doesn't have a valid markdown title.", article.Filename)
    }
  }
  xmlContent += "</channel>\n</rss>\n"
  return xmlContent
}

func WriteRSS(rssContent string, feed Feed) error {
  rssByte := []byte(rssContent)
  err := os.WriteFile(feed.OutputFile, rssByte, 0644)
  return err
}

