// Module responsible for writing RSS compliant XML based on a feed objet. ToXML function
// is called by the main control flow with a Feed object. The remaining functions add
// the header and item strings and write the XML to the file system.

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
  if isSet(article.Topic) {
    xmlContent += "\t\t<category>" + article.Topic + "</category>\n"
  }
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

func createXMLByteString(articles []Article, config Config) []byte {
  xmlContent := addHeader(config)
  for _, article := range articles {
    if len(article.Title) != 0 {
      xmlContent = addItem(xmlContent, config, article)
      Info.Printf("Added '%s' to RSS feed. ", article.Title)
    } else {
      Warn.Printf("%s doesn't have a valid markdown title.", article.Filename)
    }
  }
  xmlContent += "</channel>\n</rss>\n"
  return []byte(xmlContent)
}

func mapTopicToFile(feed *Feed) map[string][]Article {
  fileMap := make(map[string][]Article)
  for _, article := range feed.Articles {
    _, ok := fileMap[article.Topic]
    if !ok {
      if !isSet(article.Topic) {
        fileMap["index"] = []Article{article}
      } else {
        fileMap[article.Topic] = []Article{article}
      }
    } else {
      fileMap[article.Topic] = append(fileMap[article.Topic], article)
    }
  }
  return fileMap
}

func (feed *Feed) ToXML() error {
  if isSet(feed.config.OutputFolder) {
    fileMap := mapTopicToFile(feed)
    for filename, articles := range fileMap {
      rssByte := createXMLByteString(articles, feed.config)
      filepath := feed.config.OutputFolder + "/" + filename + ".xml"
      fileErr := os.WriteFile(filepath, rssByte, 0644)
      if fileErr != nil { return fileErr }
      Info.Printf("Content written to %s", filepath)
    }
  } else {
    rssByte := createXMLByteString(feed.Articles, feed.config)
    Info.Printf("Content written to %s", feed.config.OutputFile)
    return os.WriteFile(feed.config.OutputFile, rssByte, 0644)
  }
  return nil
}

