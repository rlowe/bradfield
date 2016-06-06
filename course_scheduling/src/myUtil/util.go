package myUtil

import (
  "errors"
)

func EListToAList(elist [][]string) map[string][]string {
  alist := make(map[string][]string)
  for _, a := range elist {
    l := a[0]
    r := a[1]
    if _, ok := alist[r]; !ok {
      alist[r] = []string{}
    }
    if _, ok := alist[l]; !ok {
      alist[l] = []string{}
    }

    alist[l] = append(alist[l], r)
  }
  return alist
}

// Shift removes the first element of an array and returns
// that element and modifies the array
func Shift(arr *[]string) (string, error) {
  var returnString string

  if len(*arr) == 0 {
    return returnString, errors.New("Cannot shift an empty []string")
  }

  returnString = (*arr)[0]
  *arr = (*arr)[1:len(*arr)]

  return returnString, nil
}

func InArray(needle string, haystack []string) bool {
  for _, v := range haystack {
    if needle == v {
      return true
    }
  }
  return false
}
