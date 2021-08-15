package main

import (
	"fmt"
)

func main() {
  var res int
  number := 19
  res = addDigits(number)
  fmt.Printf("result is %d:", res)
}

func addDigits(num int) int {
  sum := 0
  for num > 0 {
    sum += num % 10
	num /= 10
  } 
  if (sum >= 10) {
	return addDigits(sum)
  }
  return sum
}

// func addDigits(num int) int {
// 	return (num - 1) % 9 + 1
//   }

