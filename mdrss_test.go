package main

import (
  mdrss "mdrss/lib"
  "testing"
)

func TestFileExists(t *testing.T){
  got := mdrss.FileExists("main.go")
  want := true 
  if got != want {
    t.Errorf("got %v, wanted %v", got, want)
  } else {
    t.Log("Passed FileExists test.")
  }
}
