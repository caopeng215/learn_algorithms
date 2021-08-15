package main

import (
  "fmt"
)

func main() {
  k := 2
  var res []int
  nums := []int{ 1, 1, 1, 2, 2, 3 }
  res = topFrequent(nums, k)
  fmt.Printf("top frequent numbers is:%v", res)
}

// 方法一： 桶排序
func topFrequent(nums []int, k int) []int {
    numMap := make(map[int]int, k)
	for i := 0; i < len(nums); i++ {
	  numMap[nums[i]]++
	}
	tmps := make([]int, len(nums))
	for k, v := range numMap {
		tmps[v] = k
	}
	ret := make([]int, 0, k)
	for i := len(tmps) - 1; i > 0 && len(ret) < k; i-- {
	    if tmps[i] != 0 {
		  ret = append(ret, tmps[i])
		}
	}
	return ret
}

// 方法二：最小堆
// func topFrequent(nums []int, k int) []int {
  
// }

