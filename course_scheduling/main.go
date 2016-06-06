package main

import (
  "fmt"
  "wordladder"
  "wordladder/err"
)

func main() {
  wordList := []string {"bot","cig", "cog", "dit", "dut", "hot", "hit" ,"dot","dog","lot","log"}

  shortestPath, er := wordladder.FewestHops("hit", "aaa", wordList)
  if er != nil {
    if wle, ok := er.(*err.Error); ok {
      fmt.Printf("Error %d: %s", wle.ErrCode, wle.ErrMsg)
    }
  } else {
    fmt.Println(shortestPath)
  }
}
