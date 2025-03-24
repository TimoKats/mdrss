// Module responsible for reading markdown files and converting it to valid HTML.
// FromConfig is called by the main control flow. It adds the config first and then
// the articles using a (semi) recursive function. Remaining functions are helpers
// for converting markdown to HMTL, etc.

package lib

import (
	"bufio"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
)

func getLastFolder(path string) string {
	cleanedPath := filepath.Clean(path)
	parts := strings.Split(cleanedPath, string(filepath.Separator))
	if len(parts) == 0 {
		Error.Println("No folder found.")
		return ""
	}
	return parts[len(parts)-1]
}

func formatGuid(file fs.DirEntry) string {
	filename := strings.Split(file.Name(), ".")[0]
	return strings.ReplaceAll(filename, " ", "-")
}

func validFilename(file fs.DirEntry) bool {
	return !strings.HasPrefix(file.Name(), "draft-") && strings.HasSuffix(file.Name(), ".md")
}

func checkMarkdownTitle(text string) bool {
	if len(text) > 0 {
		return string(text[0]) == string("#")
	}
	return false
}

func convertMarkdownToXml(text []byte) string {
	opts := html.RendererOptions{Flags: html.CommonFlags}
	renderer := html.NewRenderer(opts)
	html := markdown.ToHTML(text, nil, renderer)
	return string(html)
}

func parseMarkdown(article Article, config Config) Article {
	articleBody := []byte("")
	filePath := config.InputFolder + "/" + article.Topic + "/" + article.Filename
	readFile, _ := os.Open(filePath)
	scanner := bufio.NewScanner(readFile)
	for scanner.Scan() {
		if checkMarkdownTitle(scanner.Text()) && len(article.Title) == 0 {
			article.Title = scanner.Text()[2:len(scanner.Text())]
		} else {
			articleBody = append(articleBody, scanner.Text()...)
			articleBody = append(articleBody, "\n"...)
		}
		article.Description = convertMarkdownToXml(articleBody)
	}
	return article
}

func newArticle(file fs.DirEntry, config Config) Article {
	fileInfo, _ := file.Info()
	article := Article{
		Filename:      file.Name(),
		DatePublished: fileInfo.ModTime(),
		Guid:          config.Link + "/" + formatGuid(file),
	}
	if len(config.topicInputFolder) > 0 {
		article.Topic = getLastFolder(config.topicInputFolder)
	}
	return parseMarkdown(article, config)
}


func (feed *Feed) getTopics() []string {
	var topics []string
	entries, err := os.ReadDir(feed.config.InputFolder)
	if err != nil {
		Error.Println(err)
	}
	for _, entry := range entries {
		if entry.IsDir() {
			topics = append(topics, entry.Name())
		}		
	}
	return topics
}

func getArticles(topic string, config Config) ([]Article, error) {
	var articles []Article
	rawArticles, err := os.ReadDir(config.InputFolder + "/" + topic)
	if err != nil {
		return articles, err
	}
	for _, file := range rawArticles {
		if !file.IsDir() && validFilename(file) {
			articles = append(articles, newArticle(file, config))
		}
	}
	return articles, nil
}

func (feed *Feed) FromConfig(config Config) error {
	feed.config = config
	topics := feed.getTopics()
	for _, topic := range topics {
		articles, err := getArticles(topic, feed.config)
		if err != nil {
			return err
		}
		feed.Articles = append(feed.Articles, articles...)
	}
	return nil
}
