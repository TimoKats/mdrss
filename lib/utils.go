package lib

import (
  "os"
)

func FileExists(filename string) bool {
  if _, err := os.Stat(filename); err != nil {
    return false
  }
  return true
}

func GetConfigPath() string {
  dirname, _ := os.UserHomeDir()
  return dirname + "/.mdrss/config.json"
}

