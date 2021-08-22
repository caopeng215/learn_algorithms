package main

import (
  "fmt"
)

func main() {
  nums := []int{1,2,3}
  ret := premute(nums)
  fmt.Printf("result is: %v", ret)
}

func premute(nums []int) [][]int {
  path := []int{}
  visited := map[int]bool{}
  result := [][]int{}
  helper(nums, visited, path, &result)
  return result
}

func helper(nums []int, visited map[int]bool, path []int, res *[][]int) {
  if len(path) == len(nums) {
	tmp := make([]int, len(path))
	copy(tmp, path)
	*res = append(*res, tmp)
  }
  for i:= 0; i < len(nums); i++ {
	if visited[i] {
	  continue
	}
	path = append(path, nums[i])
	visited[i] = true
	helper(nums, visited, path, res)
	visited[i] = false
	// 此处操作原因分析：
	// 在path = append(path, nums[i]) path的长度+1，
	// 然后进入helper(nums, visited, path, res)处理后，退出递归函数需还原到上一步+1前的长度
	// for循环需要在此基础上继续遍历，寻找正确的解
	path = path[:len(path)-1]
  }
}