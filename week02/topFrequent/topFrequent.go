package main

import (
  "fmt"
  "sort"
  "container/heap"
)

func main() {
  k := 2
  var res []int
  nums := []int{ 1, 1, 1, 2, 2, 3 }
  res = topFrequent(nums, k)
  fmt.Printf("top frequent numbers is:%v", res)
}

// 方法一： 桶排序
// func topFrequent(nums []int, k int) []int {
//     numMap := make(map[int]int, k)
// 	for i := 0; i < len(nums); i++ {
// 	  numMap[nums[i]]++
// 	}
// 	tmps := make([][]int, len(nums)+1)
// 	for k, v := range numMap {
// 		tmps[v] = append(tmps[v], k)
// 	}
// 	ret := make([]int, 0)
// 	for i := len(tmps) - 1; i > 0; i-- {
// 		ret = append(ret, tmps[i]...)
// 		k -= len(tmps[i])
// 		if k == 0 {
// 			break
// 		}
// 	}
// 	return ret
// }

// 方法二：最小堆
type HP struct{
  sort.IntSlice
}
func (hp HP) Less(i, j int) bool {
	return hp.IntSlice[i] < hp.IntSlice[j]
}

func (hp *HP) Push(value interface{}) {
   hp.IntSlice = append(hp.IntSlice, value.(int))
}

func (hp *HP)Pop() interface{} {
  old := hp.IntSlice
  value := old[len(old)-1]
  hp.IntSlice = old[:len(old)-1]
  return value
}

func topFrequent(nums []int, k int) []int {
  hp := &HP{make([]int, k)}
  for i:=0; i<k; i++ {
	hp.IntSlice[i] = i
  }
  heap.Init(hp)
  fmt.Printf("value is: %v", *hp)
  ret := make([]int, 0)
  top := nums[hp.IntSlice[0]]
  for i:=k; i<len(nums); i++ {
	if (nums[i] < top) {
	  heap.Push(hp, i)
	}
  }
  return []int{}
}

