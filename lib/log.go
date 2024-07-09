package lib

import (
  "log"
  "os"
)

var (
  Action *log.Logger
  Warn *log.Logger
  Info *log.Logger
  Error *log.Logger
  Fatal *log.Logger
)

var Reset = "\033[0m" 
var Red = "\033[31m"
var Yellow = "\033[33m" 
var Blue = "\033[34m" 
var Magenta = "\033[35m"

func init() {
  Info = log.New(os.Stdout, Blue + "INFO:     " + Reset, log.Ltime|log.Lshortfile)
  Warn = log.New(os.Stdout, Yellow + "WARNING:  " + Reset, log.Ltime|log.Lshortfile)
  Error = log.New(os.Stdout, Red + "ERROR:    " + Reset, log.Ltime|log.Lshortfile)
  Fatal = log.New(os.Stdout, Magenta + "FATAL:    " + Reset, log.Ltime|log.Lshortfile)
}
