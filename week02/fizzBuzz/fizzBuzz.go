package main

import( 
  "fmt"
  "strconv"
  )

func main() {
  n := 15
  res := fizzBuzz(n)
  fmt.Print("result:%v", res)
}

func fizzBuzz(n int) []string {
 res := make([]string, 0, n)
  for i := 1; i <= n; i++ {
	if i % 3 == 0 && i % 5 == 0 {
	  res = append(res, "FizzBuzz")
	} else if (i % 3 == 0) {
	  res = append(res, "Fizz")
	} else if (i % 5 == 0) {
	  res = append(res, "Buzz")
	} else {
	  res = append(res, strconv.Itoa(i))
	}
  }
  return res
}