// Module responsible for writing RSS compliant XML based on a feed objet. ToXML function
// is called by the main control flow with a Feed object. The remaining functions add
// the header and item strings and write the XML to the file system.

// Update: topics are added, so now ToXML can alternate between writing separate files
// or a single file where the topics are part of the tags.

package lib

import (
	"os"
	"time"
)

func createFileName(config Config, topic string) string {
	if topic != "" {
		return config.OutputFolder + "/" + topic + ".xml"
	}
	return config.OutputFolder + "/index.xml"
}

func addItem(xmlContent string, config Config, article Article) string {
	timestamp := article.DatePublished.Format(time.RFC822Z)
	xmlContent += "\t<item>\n"
	xmlContent += "\t\t<title>" + article.Title + "</title>\n"
	xmlContent += "\t\t<link>" + config.Link + "</link>\n"
	xmlContent += "\t\t<pubDate>" + timestamp + "</pubDate>\n"
	xmlContent += "\t\t<guid>" + article.Guid + "</guid>\n"
	if len(article.Topic) > 0 {
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

func (feed *Feed) createRssFeed(topic string) []byte {
	xmlContent := addHeader(feed.config)
	for _, article := range feed.Articles {
		if article.Topic != topic && len(feed.config.OutputFolder) > 0 {
			continue
		}
		if len(article.Title) != 0 {
			xmlContent = addItem(xmlContent, feed.config, article)
			Info.Printf("Added '%s' to RSS feed. ", article.Title)
		} else {
			Warn.Printf("%s doesn't have a valid markdown title.", article.Filename)
		}
	}
	xmlContent += "</channel>\n</rss>\n"
	return []byte(xmlContent)
}

func (feed *Feed) ToXML() error {
	if len(feed.config.OutputFolder) > 0 {
		for _, topic := range feed.topics {
			rssByte := feed.createRssFeed(topic)
			filepath := createFileName(feed.config, topic)
			fileErr := os.WriteFile(filepath, rssByte, 0644)
			if fileErr != nil {
				return fileErr
			}
			Info.Printf("Content written to %s", filepath)
		}
	} else {
		rssByte := feed.createRssFeed("")
		Info.Printf("Content written to %s", feed.config.OutputFile)
		return os.WriteFile(feed.config.OutputFile, rssByte, 0644)
	}
	return nil
}
