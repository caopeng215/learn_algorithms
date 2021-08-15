package main

import (
	"fmt"
	"sort"
	"container/heap"
)

type MinHeap struct {
  sort.IntSlice
}

func (hp MinHeap) Less(i, j int) bool {
	return hp.IntSlice[i] < hp.IntSlice[j]
}

func (hp *MinHeap) Push(value interface{}) {
	hp.IntSlice = append(hp.IntSlice, value.(int))
}

func (hp *MinHeap) Pop() interface{} {
	old := hp.IntSlice
	value := old[len(old) - 1]
	hp.IntSlice = old[:len(old) - 1]
	return value 
}

func main() {
	n := 10
	ret := nthUglyNumber(n)
	fmt.Printf("result is: %v", ret)
}

// 方法一：最小堆
var factors = []int{2, 3, 5}
func nthUglyNumber(n int) int {
  hp := &MinHeap{sort.IntSlice{1}}
  seen := map[int]struct{}{1:{}}
  for i:=1; ;i++{
    x := heap.Pop(hp).(int)
	if i == n {
	  return x
	}
	for _, v := range factors {
	  num := i * v
      if _, ok := seen[num]; !ok {
        heap.Push(hp, num)
		seen[num] = struct{}{}
	  }
	}
  }
}


// 方法二：动态规划
// func nthUglyNumber(n int) int {
//   dp := make([]int, n+1)
//   dp[1] = 1
//   a, b, c := 1, 1, 1
//   for i := 2; i <= n ; i++ {
// 	dp[i] = min(min(dp[a] * 2, dp[b] * 3), dp[c] * 5)
// 	if dp[i] == dp[a] * 2 {
// 		a++
// 	}
// 	if dp[i] == dp[b] * 3 {
// 		b++
// 	}
// 	if dp[i] == dp[c] * 5 {
// 		c++
// 	}
//   }
//   return dp[n]
// }

// func min(a, b int) int {
//   if a < b {
// 	return a
//   }
//   return b
// }