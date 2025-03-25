// various utils functions. Mainly used for filesystem related things.

package lib

import (
  "path/filepath"
  "strings"
  "io/fs"
  "os"
)

func fileExists(filename string) bool {
  if _, err := os.Stat(filename); err != nil {
    return false
  }
  return true
}

func isSet(param string) bool {
	return len(param) > 0
}

func getLastFolder(path string) string {
  cleanedPath := filepath.Clean(path)
  parts := strings.Split(cleanedPath, string(filepath.Separator))
  if len(parts) == 0 {
    Error.Println("No folder found."); return ""
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
  if isSet(text) {
    return string(text[0]) == string("#")
  }
  return false
}

