package main
// https://www.codewars.com/kata/526d84b98f428f14a60008da/

import (
  "math"
  "fmt"
)

func main() {
}

func Hammer(n int) uint {

  
  fmt.Printf("2^%v 3^%v 5^%v \n",
  math.Pow(float64(2), float64(n/2)),
  math.Pow(float64(3), float64(n/3)),
  math.Pow(float64(5), float64(n/5)),
  )

  powersOf := []int{2,3,5}

  var result float64 = 1
  for _, powerOf := range powersOf {
    var divided int = 0
    
    if (n % powerOf == 0) {
      divided = n / powerOf
    }

    fmt.Printf("div: %v\n", divided)
    // n -= divided
    res := math.Pow(float64(powerOf), float64(divided))

    fmt.Printf("res: %v\n", res)

    if(res > 1) {

    n -= int(res)
    }
    result *= res
  }

return uint(result);
}
