package main
// https://www.codewars.com/kata/526d84b98f428f14a60008da/


func main() {
}

func Hammer(n int) uint {
  if(n == 1) {
    return 1
  }
  
  var i int = 0
  var j int = 0
  var k int = 0

  ham_num := []uint{1}

  for m:= 0; m < n -1; m++ {
    var next_ham_i uint = 2 * ham_num[i]
    var next_ham_j uint = 3 * ham_num[j]
    var next_ham_k uint = 5 * ham_num[k]

    minimum := min(
      next_ham_i,
      next_ham_j,
      next_ham_k,
    )

    ham_num = append(ham_num, minimum)

    if(next_ham_i == minimum) {
      i++
    }

    if(next_ham_j == minimum) {
      j++
    }

    if(next_ham_k == minimum) {
      k++
    }
  }

  return uint(ham_num[n-1])
}

func min(a,b,c uint) uint {
  if(a < b && a < c) {
    return a
  }

  if(b < c) {
    return b
  }

  return c
}
