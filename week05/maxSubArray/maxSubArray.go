package main

import (
  "fmt"
)

func main() {
  nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
  res := maxSubArray(nums)
  fmt.Printf("result is %v", res)
}

func maxSubArray(nums []int) int {
  sum := 0
  ans := nums[1]
  var max func(a, b int) int
  max = func(a, b int) int {
	if a > b {
	  return a
	}
	return b
  }
  for _, num := range nums {
	if sum > 0 {
	  sum+=num
	} else {
	  sum = num
	}
	ans = max(sum, ans)
  }
  return ans
}