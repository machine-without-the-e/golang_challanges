package main

import (
  "testing"
)


type testcase struct {
  expected [][]int
  size int
  actual [][]int
}

func (test *testcase) execute() {
  test.actual = Spiralize(test.size)
}

func (test *testcase) isEqual() bool {
  if(len(test.expected) != len(test.actual)) {
    return false;
  }

  for row_index, row := range test.expected {
    if (len(row) != len(test.actual[row_index])) {
      return false;
    }

    for column_index, column := range row {
      if (column != test.actual[row_index][column_index]) {
        return false;
      }
    }
  }

  return true;
}

func Test_Spiralize(t *testing.T) {
  tests := []testcase{
    testcase{
      expected: [][]int{ {1,1,1,1,1},
                      	{0,0,0,0,1},
                      	{1,1,1,0,1},
                      	{1,0,0,0,1},
                      	{1,1,1,1,1}},
      size: 5,
    },
    testcase{
      expected: [][]int{ {1,1,1,1,1,1,1,1},
                      	{0,0,0,0,0,0,0,1},
                      	{1,1,1,1,1,1,0,1},
                      	{1,0,0,0,0,1,0,1},
                      	{1,0,1,0,0,1,0,1},
                      	{1,0,1,1,1,1,0,1},
                      	{1,0,0,0,0,0,0,1},
                      	{1,1,1,1,1,1,1,1}},
      size: 8,
    },
  }
  
  for _, test := range tests {
    test.execute()
    
    if (test.isEqual()) {
      continue;
    }

    t.Errorf(`Expected
    %v 
    but got
    %v
    with size %d`,
      test.expected,
      test.actual,
      test.size,
    )
  }
  // for _, test := range tests {
  //   result := Spiralize(test.size)
  //
  //   for index, row := range result {
  //     if(row == test.expected[index]) {
  //       continue;
  //     }
  //
  //     t.Errorf(`Expected
  //     %v 
  //     but got
  //     %v
  //     with size %d`,
  //       test.expected,
  //       result,
  //       test.size,
  //     )
  //   }
  //
  // }
}
