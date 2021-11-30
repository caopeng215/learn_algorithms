package main

import "fmt"

func main() {
  nums := []int{2, 7, 11, 15}
  target := 13
  fmt.Printf("result is %v:", twoSum(nums, target))
}

func twoSum(nums []int, target int) []int {
  if len(nums) < 1 {
	return []int{}
  }
  m := make(map[int]int, len(nums))
  for i, v := range nums {
	if j, ok := m[target-v]; ok {
	  return []int{j, i}
	} else {
	  m[v] = i
	}
  }
  return []int{}
}