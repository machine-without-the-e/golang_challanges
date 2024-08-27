package main

import (
  "strconv"
  "strings"
)

func DigitalRoot(n int) int {
  number := strconv.Itoa(n)

  numbers := strings.Split(number, "")

  n = 0;
  for _, numPart := range numbers {
    numPartInt, _ := strconv.Atoi(numPart);

    n += numPartInt
  }

  if (n > 9) {
    return DigitalRoot(n)
  }

  return n;
}
