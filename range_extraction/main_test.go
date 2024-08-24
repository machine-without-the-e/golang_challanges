package main

import (
  "testing"
)

type testcase struct {
  expected string
  list []int
}


func Test_Solution(t *testing.T) {
  
  tests := []testcase{
    testcase{
      expected: "-10",
      list: []int{-10},
    },
    testcase{
      expected: "-10--8",
      list: []int{-10, -9, -8},
    },
    testcase{
      expected: "-10--8,-6,-3-1,3-5,7-11,14,15,17-20",
      list: []int{-10, -9, -8, -6, -3, -2, -1, 0, 1, 3, 4, 5, 7, 8, 9, 10, 11, 14, 15, 17, 18, 19, 20},
    },
    testcase{
      expected: "-6,-3-1,3-5,7-11,14,15,17-20",
      list: []int{-6,-3,-2,-1,0,1,3,4,5,7,8,9,10,11,14,15,17,18,19,20},
    },
  }

  for _, test := range tests {
    result := Solution(test.list)

    if(result == test.expected) {
      continue;
    }

    t.Errorf("Expected `%s`, but got `%s` - with input `%v`",
      test.expected,
      result,
      test.list,
    )
  }

}
