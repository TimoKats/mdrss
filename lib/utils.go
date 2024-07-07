package lib

import (
	"bufio"
	"fmt"
  "os"
)

var Cyan = "\033[36m" 

func ConsoleInput(message string) string {
	in := bufio.NewReader(os.Stdin)
	fmt.Print(Cyan + "ACTION:   " + Reset + message)
	command, _ := in.ReadString('\n')
	return string(command)
}

func FileExists(filename string) bool {
  if _, err := os.Stat(filename); err != nil {
    return false
  }
  return true
}
