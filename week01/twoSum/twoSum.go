package main

import "fmt"

func main() {
  nums := []int{1,2,3,4,5,6,7}
  k := 10
  arr := twoSum(nums, k)
  fmt.Printf("exist: %v", arr)
}

func twoSum(nums []int, target int) []int {
  numMap := make(map[int]int, len(nums))
  for i:=0; i < len(nums); i++ {
      numMap[nums[i]] = i
  }
  for i:=0; i<len(nums); i++ {
      if index, ok := numMap[target - nums[i]]; ok && i != index {
         return []int {i, index} 
      }
  }
  numMap = nil
  return []int{}
}
