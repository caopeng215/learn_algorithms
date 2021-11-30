package main

import "fmt"

func main() {
  arrs := []int{0, 1, 0, 3, 12}
  moveZeroes(arrs)
}

func moveZeroes(nums []int) {
  j := 0
  for i := 0; i < len(nums); i++ {
	if nums[i] != 0 {
	  tmp := nums[i]
	  nums[i] = nums[j]
	  nums[j] = tmp
	  j++
	}
  }
  fmt.Printf("ans is: %v", nums)
}