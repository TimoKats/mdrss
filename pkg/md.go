// Module responsible for reading markdown files and converting it to valid HTML.
// FromConfig is called by the main control flow. It adds the config first and then
// the articles using a (semi) recursive function. Remaining functions are helpers
// for converting markdown to HMTL, etc.

package lib

import (
	"bufio"
	"io/fs"
	"os"
	"strings"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"	
)

func formatGuid(file fs.DirEntry) string {
	filename := strings.Split(file.Name(), ".")[0]
	return strings.ReplaceAll(filename, " ", "-")
}

func validFilename(file fs.DirEntry) bool {
	return !strings.HasPrefix(file.Name(), "draft-") && 
	strings.HasSuffix(file.Name(), ".md")
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

func newArticle(config Config, file fs.DirEntry, topic string) Article {
	fileInfo, _ := file.Info()
	article := Article{
		Filename:      file.Name(),
		DatePublished: fileInfo.ModTime(),
		Guid:          config.Link + "/" + formatGuid(file),
		Topic:         topic,
	}
	return parseMarkdown(article, config)
}

func (feed *Feed) getArticles(topic string) ([]Article, error) {
	var articles []Article
	rawArticles, err := os.ReadDir(feed.config.InputFolder + "/" + topic)
	if err != nil {
		return articles, err
	}
	for _, file := range rawArticles {
		if !file.IsDir() && validFilename(file) {
			articles = append(articles, newArticle(feed.config, file, topic))
		}
	}
	return articles, nil
}

func (feed *Feed) getTopics() error {
	entries, err := os.ReadDir(feed.config.InputFolder)
	if err != nil {
		return err
	}
	for _, entry := range entries {
		if entry.IsDir() {
			feed.topics = append(feed.topics, entry.Name())
		}
	}
	feed.topics = append(feed.topics, "") // for root folder articles
	return nil
}

func (feed *Feed) FromConfig(config Config) error {
	feed.config = config
	if 	topicErr := feed.getTopics(); topicErr != nil {
		return topicErr
	}
	for _, topic := range feed.topics {
		articles, articleErr := feed.getArticles(topic)
		if articleErr != nil {
			return articleErr
		}
		feed.Articles = append(feed.Articles, articles...)
	}
	return nil
}
