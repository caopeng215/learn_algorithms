package main

import (
	"container/heap"
	"sort"
	"fmt"
)

var a []int

type PriorityQueue struct {
  sort.IntSlice
}

func (pq PriorityQueue) Less(i, j int) bool {
	return a[pq.IntSlice[i]] > a[pq.IntSlice[j]]
}

func(pq *PriorityQueue) Push(value interface{}) {
   pq.IntSlice = append(pq.IntSlice, value.(int))
//    pq.IntSlice[len(pq.IntSlice)-1] = value.(int) // 等价
}

func (pq *PriorityQueue) Pop() interface{} {
	old := pq.IntSlice
	value := old[len(old)-1]
	pq.IntSlice = old[:len(old)-1]
	return value
}

func main() {
  nums := []int{1,3,-1,-3,5,3,6,7}
  a = nums
  ret := maxSlidingWindow(nums, 3)
  fmt.Printf("result is: %v", ret)
  //  [3,3,5,5,6,7] 结果
}

func maxSlidingWindow(nums []int, k int) []int {
  if len(nums) == 0 || k == 0 {
	return []int{}
  }
  pq := &PriorityQueue{make([]int, k)}
  for i:=0;i<k;i++{
	pq.IntSlice[i] = i
  }
  heap.Init(pq)
  result := make([]int, 1, len(nums)-k+1)
  result[0] = a[pq.IntSlice[0]]
  for i:=k;i<len(nums);i++{
	heap.Push(pq,i)
	for pq.IntSlice[0] <= i-k {
      heap.Pop(pq)
	}
	result = append(result, nums[pq.IntSlice[0]])
	// result[i-k+1] = nums[pq.IntSlice[0]] //等价
  }
  return result
}