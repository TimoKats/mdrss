// logging class, created upon init.

package lib

import (
	"log"
	"os"
)

var (
	Action *log.Logger
	Warn   *log.Logger
	Info   *log.Logger
	Error  *log.Logger
	Fatal  *log.Logger
)

var reset = "\033[0m"
var red = "\033[31m"
var yellow = "\033[33m"
var blue = "\033[34m"
var magenta = "\033[35m"

func init() {
	Info = log.New(os.Stdout, blue+"info:     "+reset, log.Ltime)
	Warn = log.New(os.Stdout, yellow+"warning:  "+reset, log.Ltime)
	Error = log.New(os.Stdout, red+"error:    "+reset, log.Ltime)
	Fatal = log.New(os.Stdout, magenta+"fatal:    "+reset, log.Ltime)
}
