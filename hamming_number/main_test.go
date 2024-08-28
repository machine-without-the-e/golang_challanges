package main

import (
  "testing"
)

func dotest(input int, expected uint, t *testing.T) {
  result := Hammer(input)

  if(result == expected) {
  
    return
  }

  t.Errorf("Expected %d, got %d, with input %d",
    expected,
    result,
    input,
  )
}

func Test_hammingnumber(t *testing.T) {
  // dotest(1, 1, t)
  // dotest(2, 2, t)
  // dotest(3, 3, t)
  // dotest(4, 4, t)
  // dotest(5, 5, t)
  dotest(6, 6, t)
  // dotest(7, 8, t)
  // dotest(8, 9, t)
  // dotest(9, 10, t)
  // dotest(10, 12, t)
  // dotest(11, 15, t)
  // dotest(12, 16, t)
  // dotest(13, 18, t)
  // dotest(14, 20, t)
  // dotest(15, 24, t)
  // dotest(16, 25, t)
  // dotest(17, 27, t)
  // dotest(18, 30, t)
  // dotest(19, 32, t)
}
