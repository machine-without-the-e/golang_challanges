package main

import "fmt"
//https://www.codewars.com/kata/586f6741c66d18c22800010a/go
//function (begin,end,step)

// begin > end = 0
// end is in the steps,add it to sum.

func get_last_element(arr []int) int {
  return arr[len(arr)-1]
}

func sequence (begin int, end int, step int) int {
  if (begin > end) {
    return 0;
  }
  
  var steps = []int{begin};
  
  for (get_last_element(steps) + step) <= end {
    steps = append(steps, (get_last_element(steps) + step));
  }

  sum := 0;
  for _, seq := range steps {
    sum += seq
  }

  print(fmt.Sprintf(`%d
    `, steps));

  return sum;
}

    

func main() {

  // begin, end, step, expected
  sequences := [][]int{
    {2,2,2, 2},
    {2,24,22, 26},
    {2,6,2, 12},
    {1,5,1, 15},
    {1,5,3, 5},
  }

  for _, seq := range sequences {
    result := sequence(seq[0], seq[1], seq[2]);


    success_icon := "[FAIL]";

    if (result == seq[3]) {
      success_icon = "✓";
    }

    print(fmt.Sprintf("%s %d,%d,%d: %d - should be %d",
      success_icon,
      seq[0],
      seq[1],
      seq[2],
      result,
      seq[3],
    ), "\n");
  }

  print("\n");
}
