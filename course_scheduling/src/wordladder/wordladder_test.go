package wordladder

// Tease out the data into JSON

import (
  "testing"
  "wordladder/err"
)

type tests struct {
  BeginWord string
  EndWord   string
  WordList  []string
  Hops      int
  ErrCode   int
}

var testPairs = []tests {
  {
    "hit",
    "cog",
    []string{"hot","dot","dog","lot","log"},
    5,
    0,
  },
  {
    "hit",
    "cog",
    []string{"bot","cig", "cog", "dit", "dut", "hot", "hit" ,"dot","dog","lot","log"},
    5,
    0,
  },
  {
    "hit",
    "hit",
    []string{"dot","dog","lot","log"},
    1,
    err.ALREADY_THERE,
  },
  {
    "hit",
    "aaa",
    []string{"dot","dog","lot","log"},
ZZ    1,
    err.NO_WAY,
  },
}

func TestFewestHops(t *testing.T) {
  for _, test := range testPairs {
    result, er := FewestHops(test.BeginWord, test.EndWord, test.WordList)
    if er != nil {
      if wle, ok := er.(*err.Error); ok {
        if test.ErrCode != wle.ErrCode {
          t.Error("FAILED FOR", test.BeginWord, "and", test.EndWord,
                  "EXPECTED ERR", test.ErrCode, "but got", wle.ErrCode)
        }
      }
    } else {
      if len(result) != test.Hops {
        t.Error("FAIL")
      }
    }
  }
}
