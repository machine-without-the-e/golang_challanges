package main

import (
  "testing"
)

type testdata struct {
  input int
  expected int
}

func Test_DigitalRoot(t *testing.T) {
  tests := []testdata{
    {
      input: 16,
      expected: 7,
    },
    {
      input: 942,
      expected: 6,
    },
    {
      input: 132189,
      expected: 6,
    },
    {
      input: 493193,
      expected: 2,
    },
  }
  
  for _, test := range tests {
    result := DigitalRoot(test.input)

    if(result == test.expected) {
      continue;
    }
  
    t.Errorf("Expected %d but got %d for input of %d",
      test.expected,
      result,
      test.input,
    )

  }
}
