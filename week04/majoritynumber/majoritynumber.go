package main

import "fmt"

func main() {
	nums := []int{3,2,3}
	ret := majorityElement(nums)
	fmt.Printf("result is: %v", ret)
}
func majorityElement(nums []int) int {
	numMap := map[int]int{}
	for _, num := range nums {
	  if _, ok := numMap[num]; ok {
		numMap[num]++
	  } else {
		numMap[num] = 1
	  }
	}
	fmt.Printf("numMap is: %v", numMap)
	for key, val := range numMap {
	  if val > len(nums) / 2 {
		return key
	  }
	}
	return -1
  }
