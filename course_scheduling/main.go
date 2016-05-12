package main

import (
  "course"
  "fmt"
)

func main() {
  courseList := [][]int {
    {1,2},
    {2,3},
    {2,4},
    {4,5},
    //{5,4},
  }

  order, err := course.TopOrder(courseList)
  if err != nil {
    fmt.Println(err)
    return
  }

  fmt.Println("ORDER:",order)
}
