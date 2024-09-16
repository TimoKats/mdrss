package lib

import (
  "regexp"
)

func ConvertLink(text string, markdownLinks *regexp.Regexp) string {
  return "<p>" + markdownLinks.ReplaceAllString(text, "<a href='$2'>$1</a>") + "</p>"
}

func ConvertEnclosure(text string, markdownLinks *regexp.Regexp) string {
  url := markdownLinks.ReplaceAllString(text, "$2")
  size, fileSizeErr := FileSizeUrl(url)
  if fileSizeErr != nil {
    Error.Println(fileSizeErr)
    return ""
  }
  return "<enclosure " + markdownLinks.ReplaceAllString(text, "url='$2' type='$1' length='") + size + "' />"
}

func ConvertUnorderedlList(text string) string {
  if !markdownUlListActive {
    markdownUlListActive = true
    return "<ul><li>" + text[getLeadingWhitespace(text):] + "</li>"
  }
  return "<li>" + text[getLeadingWhitespace(text):] + "</li>"
}

func ConvertOrderedlList(text string, index int64) string {
  if markdownOlIndex == 0 {
    markdownOlIndex = 1
    return "<ol><li>" + text[getLeadingWhitespace(text):] + "</li>"
  }
  return "<li>" + text[getLeadingWhitespace(text):] + "</li>"
}

func ConvertTextEnrichment(text string) string {
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
  return text
}

