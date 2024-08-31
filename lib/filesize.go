package lib

import (
  "unicode"
  "strings"
  "os/exec"
  "errors"
)

func extractInt(text string) string {
  filteredString := ""
  for _, character := range text {
    if unicode.IsDigit(character) {
      filteredString += string(character)
    }
  }
  if len(filteredString) == 0 {
    return "1"
  } else {
    return filteredString
  }
}

func FileSizeUrl(url string) (string, error) {
  cmd := exec.Command("curl", "-sIL", url)
  cmdOutput, cmdErr := cmd.Output()
  if cmdErr != nil {
    return "1", errors.New("Issues when running curl to get filesize.")
  } else {
    for _, line := range strings.Split(string(cmdOutput), "\n") {
      if strings.Contains(line, "content-length") {
        return extractInt(line), nil
      }
    }
  }
  return "1", errors.New("No content-length found for url")
}
