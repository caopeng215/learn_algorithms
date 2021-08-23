package main

import "fmt"

func main() {
  nums := []int{1, 2, 3}
  ret := permute(nums)
  fmt.Printf("result is: %v", ret)
}

func permute(nums []int) [][]int { 
  list := []int{}
  visited := map[int]bool{}
  res := [][]int{}
  helper(nums, list, visited, &res)
  return res
}

func helper(nums []int, list []int, visited map[int]bool, res *[][]int) {
  if len(nums) == len(list) {
	tmp := make([]int, len(nums))
	copy(tmp, list)
	*res = append(*res, tmp)
	return
  }
  for i := 0; i < len(nums); i++ {
	if visited[i] {
	  continue
	}
	list = append(list, nums[i])
	visited[i] = true
	helper(nums, list, visited, res)
	visited[i] = false
	list = list[:len(list)-1]
  }
}