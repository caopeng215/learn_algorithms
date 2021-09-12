package main

import "fmt"

func main() {
  var n = 2
  fmt.Printf("ans is %v", checkRecord(n))
}

func checkRecord(n int) int {
  if n == 1 {
	return 3
  }
  return 0
}