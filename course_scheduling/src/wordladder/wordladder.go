package wordladder

// https://leetcode.com/problems/word-ladder/

import (
  "myUtil"
  "wordladder/err"
)

func GenerateAnAList(wordList []string) map[string][]string {
  var aList = make(map[string][]string)

  // Generate the keys
  for _, word := range wordList {
    if _, ok := aList[word]; !ok {
      aList[word] = []string{}
    }
  }

  // Cycle through the keys and add every node that has a direct
  // relationship (two characters in common)
  for key, _ := range aList {
    a := []byte(key)
    for _, word := range wordList {
      b := []byte(word)
      // Gotta match two
      if (a[0] == b[0] && a[1] == b[1]) ||
         (a[0] == b[0] && a[2] == b[2]) ||
         (a[1] == b[1] && a[2] == b[2]) {
        aList[key] = append(aList[key], word)
      }
    }
  }
  return aList
}

func FewestHops(beginWord string, endWord string, wordList []string) ([]string, error) {
  wordList = append(wordList, beginWord)
  wordList = append(wordList, endWord)

  aList := GenerateAnAList(wordList)
  var queueToExplore = []string{ beginWord }
  var shortestPath []string // This is what we shall return
  var relationships = make(map[string]string)

  if beginWord == endWord {
    return shortestPath, &err.Error{err.ALREADY_THERE, "You're already there"}
  }

  for len(queueToExplore) > 0 {
    node, _ := myUtil.Shift(&queueToExplore)
    for _, neighbor := range aList[node] {
      // Ignore it if we've already done it
      if _, ok := relationships[neighbor]; !ok {
        if neighbor == endWord {
          relationships[neighbor] = node
          shortestPath = BuildPathFromRelationships(beginWord, endWord, relationships)
          return shortestPath, nil
        } else {
          relationships[neighbor] = node
          queueToExplore = append(queueToExplore, neighbor)
        }
      }
    }
  }
  return shortestPath, &err.Error{err.NO_WAY, "You can't get there from here"}
}

func BuildPathFromRelationships(beginWord string, endWord string, relationships map[string]string) []string {
  var shortestPath []string
  currentNode := endWord
  for {
    for k, v := range relationships {
      if currentNode == beginWord {
        shortestPath = append([]string{currentNode}, shortestPath...)
        return shortestPath
      } else if k == currentNode {
        shortestPath = append([]string{currentNode}, shortestPath...)
        currentNode = v
      }
    }
  }
  return shortestPath
}
