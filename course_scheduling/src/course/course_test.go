package course

import (
  "course"
  "fmt"
  "reflect"
  "testing"
)

type eListtoAlistTests struct {
  Given    [][]int       `json:"given"`
  Expected map[int][]int `json: "expected"`
}

var eListtoAlistExamples = []eListtoAlistTests {
  eListtoAlistTests {
    [][]int{ { 1, 2 }, { 2, 3 } },
    map[int][]int{ 1: { 2 }, 2: { 3 }, 3: {} },
  },
}

func TestElistToAlist(T *testing.T) {
  for _, test := range eListtoAlistExamples {
    actual := course.EListToAList(test.Given)
    if !reflect.DeepEqual(actual, test.Expected) {
      fmt.Println("[FAIL] Expected ",
                  test.Expected,
                  " for ", test.Given,
                  " but got ", actual)
      T.Fail()
    } else {
      fmt.Print(".")
    }
  }
}

type DepthFirstWalkTests struct {
  GivenNode            int `json:"given_node"`
  GivenGraph map[int][]int `json:"given_graph"`
  ExpectedVal        []int `json:"expected_val"`
  ExpectedError      error `json:"expected_error"`
}

var DepthFirstWalkExamples = []DepthFirstWalkTests {
  DepthFirstWalkTests {
    1,
    map[int][]int{ 1: {} },
    []int{ 1 },
    nil,
  },
  DepthFirstWalkTests {
    1,
    map[int][]int{ 1: { 2 }, 2: { } },
    []int{ 1, 2 },
    nil,
  },
}

func TestDepthFirstWalk(T *testing.T) {
  for _, test := range DepthFirstWalkExamples {
    actual, err := course.DepthFirstWalk(test.GivenGraph, test.GivenNode)

    if !(reflect.DeepEqual(actual, test.ExpectedVal) && err == test.ExpectedError) {
      fmt.Println("[FAIL] Expected Val ", test.ExpectedVal, " Expected Err ", 
                  test.ExpectedError, "for ", test.GivenNode, "and ", test.GivenGraph,
                  " but got ", actual, " and ", err)
      T.Fail()
    } else {
      fmt.Print(".")
    }
  }
}
