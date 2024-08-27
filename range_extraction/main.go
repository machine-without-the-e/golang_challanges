package main

import (
  "strconv"
  "strings"
)

func Solution(list []int) string {

  result := []string{}
  buffer := []int{}

  for _, number := range list {
    if (len(buffer) == 0) {
      buffer = append(buffer, number)
      continue
    }
    
    last_number := get_last_element(buffer)

    if (last_number + 1 == number) {
      buffer = append(buffer, number)
      continue
    }

    // Range halted
    result, buffer = appendBuffer(number, result, buffer)
  }

  result, buffer = appendBuffer(0, result, buffer)
  return strings.Join(result, ",")
}

func appendBuffer(number int, result []string, buffer []int) ([]string, []int) {

    last_number := get_last_element(buffer)
    if (len(buffer) >= 3) {
      result = append(result, strconv.Itoa(buffer[0]) + "-" + strconv.Itoa(last_number))
      buffer = []int{numbere

      return result, buffer
    }

    for _, bufferNumber := range buffer {
      result = append(result, strconv.Itoa(bufferNumber))
    }
    
    buffer = []int{number}

    return result, buffer
}

func get_last_element(arr []int) int {
  return arr[len(arr)-1]
}

func main() {

}
