package main

import (
  "fmt"
)

func main() {
  var target = 6
  nums := []int{3,2,4}
  ret := twoSum(nums, target)
  fmt.Printf("result is: %v", ret)
}

func twoSum(nums []int, target int) []int {
  numMap := map[int]int{}
  for i, v := range nums {
	if j, ok := numMap[target-v]; ok && i != j {
	  return []int{j, i}
	} else {
		numMap[v] = i
	}
  }
  return []int{}
}