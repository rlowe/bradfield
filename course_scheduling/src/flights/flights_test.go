package flights

import (
  "errors"
  "fmt"
  "testing"
)

var airports = [][]string {
  { "lax", "foo" },
  { "lax", "ord" },
  { "sfo", "iah" },
  { "foo", "iah" },
  { "ord", "iah" },
  { "iah", "lga" },
  { "iah", "mia" },
  { "ord", "mia" },
}

type testpair struct {
  origin      string
  destination string
  paths       [][]string
  expectedHops int
  err          error
}

var tests = []testpair {
  { "lax", "mia", airports, 3, nil},
  { "lax", "lax", airports, 0, errors.New("You're already there")},
  { "lax", "nyc", airports, 0, errors.New("You can't get there from here")},
}

func TestFewestHops(t *testing.T) {
  for _, pair := range tests {
    result, _ := FewestHops(pair.origin, pair.destination, pair.paths)
    if len(result) != pair.expectedHops {
      fmt.Println(result)
      t.Error("FAIL")
    }
  }
}
