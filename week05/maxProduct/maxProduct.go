package main

import "fmt"

func main() {
  nums := []int{-3, -1, -1}
  fmt.Printf("result is %v:", maxProduct(nums))
}

func maxProduct(nums []int) int {
  m := len(nums)
  if m == 1 {
	return nums[0]
  }
  Max, imax, imin := nums[0], nums[0], nums[0]
  for i:=1; i<len(nums); i++ {
	if nums[i] < 0 {
	  imax, imin = imin, imax
	}
	imax = max(imax * nums[i], nums[i])
	imin = min(imin * nums[i], nums[i])
	Max = max(Max, imax)
  }
  return Max
}
func min(a, b int) int {
	if a < b {
	  return a
	}
	return b
  }
func max(a, b int) int {
  if a > b {
	return a
  }
  return b
}