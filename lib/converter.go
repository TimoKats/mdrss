package lib

import (
  "strings"
  "regexp"
)

var markdownUlListActive bool
var markdownOlIndex int64
var codeBlockActive = false
var codeBlockAggregate = ""

func ConvertReferenceLink(text string, markdownLinks *regexp.Regexp) string {
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

func ConvertUnorderedlList(text string, markdownUnorderedLists *regexp.Regexp) string {
  if markdownUnorderedLists.MatchString(text) {
    if !markdownUlListActive {
      markdownUlListActive = true
      return "<ul><li>" + text[getLeadingWhitespace(text):] + "</li>"
    }
    return "<li>" + text[getLeadingWhitespace(text):] + "</li>"
  }
  markdownUlListActive = false
  return "</ul><p>" + text + "</p>"
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

func ConvertLink(text string, markdownLinks *regexp.Regexp) string {
  if strings.Contains(text, "audio/mpeg") {
    return ConvertEnclosure(text, markdownLinks)
  } else {
    return ConvertReferenceLink(text, markdownLinks)
  }
}

func ConvertCodeblock(text string, fencedCodeBlock *regexp.Regexp) string {
  if !codeBlockActive {
    codeBlockActive = true
    codeBlockAggregate = "<sup>" + text[3:] + "</sup><br>"
    return "" + "<pre style=\"word-wrap: break-word;\"><code>"
  } else if fencedCodeBlock.MatchString(text) {
    out := codeBlockAggregate
    codeBlockAggregate, codeBlockActive = "", false
    return out + "</code></pre>"
  }
  if codeBlockAggregate != "" {
    codeBlockAggregate += "<br>"
  }
  codeBlockAggregate += text
  return ""
}

func ConvertOrderedlList(text string) string {
  if markdownOlIndex == 0 {
    markdownOlIndex = 1
    return "<ol><li>" + text[getLeadingWhitespace(text):] + "</li>"
  }
  return "<li>" + text[getLeadingWhitespace(text):] + "</li>"
}

func ConvertOrderedLists(text string, markdownOrderedLists *regexp.Regexp) string {
  if markdownOrderedLists.MatchString(text) {
    entryText := markdownOrderedLists.FindStringSubmatch(text)[4]
    return ConvertOrderedlList(entryText)
  }
  markdownOlIndex = 0
  return "</ol>" + ConvertMarkdownToRSS(text)
}

