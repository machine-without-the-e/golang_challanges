package main

import (
	"testing"
  "sort"
)

type testcase struct {
  numberOfPairs int
  expectedPairs []string
}

func Test_BalancedParens(t *testing.T) {
  tests := []testcase{
    {
      numberOfPairs: 0,
      expectedPairs: []string{},
    },
    {
      numberOfPairs: 1,
      expectedPairs: []string{"()"},
    },
    {
      numberOfPairs: 2,
      expectedPairs: []string{"(())","()()"},
    },
    {
      numberOfPairs: 3,
      expectedPairs: []string{"((()))","(()())","(())()","()(())","()()()"},
    },
    {
      numberOfPairs: 4,
      expectedPairs: []string{"(((())))","((()()))","((())())","((()))()","(()(()))","(()()())","(()())()","(())(())","(())()()","()((()))","()(()())","()(())()","()()(())","()()()()"},
    },
  }
  var sorted=func(a []string)[]string{ sort.Strings(a);return a; }
  for _, test := range tests {
    result := sorted(BalancedParens(test.numberOfPairs))

    if (len(result) != len(test.expectedPairs)) {
      t.Errorf("The result did not match expected. For `n`=%d - Expected: %v, but got %v",
        test.numberOfPairs,
        test.expectedPairs,
        result,
      )
    }

    for index, expectedPair := range test.expectedPairs {
      if(result[index] == expectedPair) {
        continue
      }

      t.Errorf("The result did not match expected. For `n`=%d - Expected: %v, but got %v",
        test.numberOfPairs,
        test.expectedPairs,
        result,
      )
    }
  }
}
