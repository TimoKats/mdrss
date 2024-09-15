package lib

import (
	"bufio"
	"errors"
	"os"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

var markdownUlListActive bool
var markdownOlIndex int64
var codeBlockAggregate, codeBlockOpen = "", false

func checkMarkdownTitle(text string) bool {
	if len(text) > 0 {
		return string(text[0]) == string("#")
	}
	return false
}

func getLeadingWhitespace(text string) int {
	for index, letter := range text {
		if unicode.IsLetter(letter) {
			return index
		}
	}
	return 0
}

func convertMarkdownLink(text string, markdownLinks *regexp.Regexp) string {
	return "<p>" + markdownLinks.ReplaceAllString(text, "<a href='$2'>$1</a>") + "</p>"
}

func convertMarkdownEnclosure(text string, markdownLinks *regexp.Regexp) string {
	url := markdownLinks.ReplaceAllString(text, "$2")
	size, fileSizeErr := FileSizeUrl(url)
	if fileSizeErr != nil {
		Error.Println(fileSizeErr)
		return ""
	}
	return "<enclosure " + markdownLinks.ReplaceAllString(text, "url='$2' type='$1' length='") + size + "' />"
}

func convertMarkdownUlList(text string) string {
	if !markdownUlListActive {
		markdownUlListActive = true
		return "<ul><li>" + text[getLeadingWhitespace(text):] + "</li>"
	}
	return "<li>" + text[getLeadingWhitespace(text):] + "</li>"
}

func convertMarkdownOlList(text string, index int64) string {
	if markdownOlIndex == 0 {
		markdownOlIndex = 1
		return "<ol><li>" + text[getLeadingWhitespace(text):] + "</li>"
	}
	return "<li>" + text[getLeadingWhitespace(text):] + "</li>"
}

func ConvertMarkdownToRSS(text string) string {
	markdownLinks := regexp.MustCompile(`\[(.*)\]\((.*)\)`)
	markdownUnorderedLists := regexp.MustCompile(`^(\s*)(-|\*|\+)[\s](.*)`)
	markdownOrderedLists := regexp.MustCompile(`^(\s*)(-?\d+)(\.\s+)(.*)$`)
	fencedCodeBlock := regexp.MustCompile("^\x60\x60\x60")

	// Strikethrough
	text = inlineRewrap(text, regexp.MustCompile(`(~~)([^~~]+)(~~)`), "<s>", "</s>")
	// Subscript
	text = inlineRewrap(text, regexp.MustCompile(`(~)([^~]+)(~)`), "<sub>", "</sub>")
	// Superscript
	text = inlineRewrap(text, regexp.MustCompile(`(\^)([^\^]+)(\^)`), "<sup>", "</sup>")

	// Strong Emphasis
	text = inlineRewrap(text, regexp.MustCompile(`(\*\*\*|___)([^\*^_]+)(\*\*\*|___)`), "<b><i>", "</b></i>")
	// Bold
	text = inlineRewrap(text, regexp.MustCompile(`(\*\*|__)([^\*^_]+)(\*\*|__)`), "<b>", "</b>")
	// Italic
	text = inlineRewrap(text, regexp.MustCompile(`(\*|_)([^\*^_]+)(\*|_)`), "<i>", "</i>")

	// Inline Codeblock
	text = inlineRewrap(text, regexp.MustCompile(`(\x60)([^\x60]+)(\x60)`), "<code>", "</code>")

	switch {
	case codeBlockOpen && !fencedCodeBlock.MatchString(text):
		if codeBlockAggregate != "" {
			codeBlockAggregate += "<br>"
		}
		codeBlockAggregate += text
		return ""
	case markdownLinks.MatchString(text):
		if strings.Contains(text, "audio/mpeg") {
			return convertMarkdownEnclosure(text, markdownLinks)
		} else {
			return convertMarkdownLink(text, markdownLinks)
		}
	case markdownUnorderedLists.MatchString(text):
		return convertMarkdownUlList(text)
	case markdownUlListActive:
		markdownUlListActive = false
		return "</ul><p>" + text + "</p>"
	case markdownOrderedLists.MatchString(text):
		entryIndex, entryErr := strconv.ParseInt(markdownOrderedLists.FindStringSubmatch(text)[2], 10, 64)
		entryText := markdownOrderedLists.FindStringSubmatch(text)[4]
		if entryErr != nil {
			return "<p>" + text + "</p>"
		}
		return convertMarkdownOlList(entryText, entryIndex)
	case markdownOlIndex != 0:
		markdownOlIndex = 0
		return "</ol>" + ConvertMarkdownToRSS(text)
	case fencedCodeBlock.MatchString(text):
		if !codeBlockOpen {
			codeBlockOpen = true
			codeBlockAggregate = "<sup>" + text[3:] + "</sup><br>"
			return "" + "<pre style=\"word-wrap: break-word;\"><code>"
		} else {
			out := codeBlockAggregate
			codeBlockAggregate, codeBlockOpen = "", false
			return out + "</code></pre>"
		}
	default:
		return "<p>" + text + "</p>"
	}
}

func inlineRewrap(text string, pattern *regexp.Regexp, prefix string, postfix string) string {
	if pattern.Match([]byte(text)) {
		out := pattern.ReplaceAllFunc([]byte(text), func(b []byte) []byte {
			return []byte(prefix + pattern.FindStringSubmatch(string(b))[2] + postfix)
		})
		return string(out)
	} else {
		return text
	}
}

func GetArticles(config Config) ([]Article, error) {
	RawArticles, fileErr := os.ReadDir(config.InputFolder)
	articles := []Article{}
	if fileErr == nil {
		for _, file := range RawArticles {
			if !file.IsDir() && !strings.HasPrefix(file.Name(), "draft-") && strings.HasSuffix(file.Name(), ".md") {
				var article Article
				fileInfo, _ := file.Info()
				article.Filename = file.Name()
				article.DatePublished = fileInfo.ModTime()
				articles = append(articles, article)
			}
		}
		return articles, nil
	}
	return articles, errors.New("Error when getting files from input directory.")
}

func ReadMarkdown(config Config, articles []Article) []Article {
	for index := range articles {
		markdownUlListActive = false
		filePath := config.InputFolder + "/" + articles[index].Filename
		readFile, _ := os.Open(filePath)
		scanner := bufio.NewScanner(readFile)
		for scanner.Scan() {
			if checkMarkdownTitle(scanner.Text()) && len(articles[index].Title) == 0 {
				articles[index].Title = scanner.Text()[2:len(scanner.Text())]
			} else if len(scanner.Text()) > 0 || codeBlockOpen {
				articles[index].Description += ConvertMarkdownToRSS(scanner.Text())
			}
		}
	}
	return articles
}
