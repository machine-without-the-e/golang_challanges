package main

import (
  "sort"
  "fmt"
)

// Given an array of integers, find the one that appears an odd number of times.
//
// There will always be only one integer that appears an odd number of times.

func FindOdd(seq []int) int {
  sort.Slice(seq, func(i, j int) bool {return seq[i] < seq[j]});
print(fmt.Sprintf("%d", seq));
  prev := seq[0];
  count := 0;
  for _, element := range seq {
    if(element == prev) {
      count++;
      continue;
    }

    if(count % 2 == 1) {
      return prev;
    }

    count = 1;
    prev = element;
  } 

  return prev;
}

func main() {

print(FindOdd([]int{20, 1, -1, 2, -2, 3, 3, 5, 5, 1, 2, 4, 20, 4, -1, -2, 5}))
  print(FindOdd([]int{1,0,1,0,1,0,2,1,0,2,2,2,2,22,2,2}))
}
