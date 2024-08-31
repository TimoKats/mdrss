package lib

import (
  "strconv"
  "unicode"
  "strings"
  "os/exec"
  "os"
)

func extractInt(text string) int {
  filteredString := ""
  for _, character := range text {
    if unicode.IsDigit(character) {
      filteredString += string(character)
    }
  }
  integer, convErr := strconv.Atoi(filteredString)
  if convErr != nil {
    Error.Println(convErr)
    return 1
  } else {
    return integer
  }
}

func FileSizeLocal(filename string) int {
  fileInfo, fileErr := os.Stat(filename)
  if fileErr != nil {
    Error.Println(fileErr)
    return 1
  }
  return int(fileInfo.Size())
}

func FileSizeUrl(url string) int {
  cmd := exec.Command("curl", "-sIL", url)
  cmdOutput, cmdErr := cmd.Output()
  if cmdErr != nil {
    Error.Println(cmdErr)
    return 1
  } else {
    for _, line := range strings.Split(string(cmdOutput), "\n") {
      if strings.Contains(line, "content-length") {
        return extractInt(line)
      }
    }
  }
  Error.Printf("No content-length found for %s", url)
  return 1
}
