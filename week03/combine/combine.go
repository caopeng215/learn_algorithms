package main

import (
	"fmt"
)

func main() {
	n, k := 4, 2
	ret := combine(n, k)
	fmt.Printf("ret is %v:", ret)
}

func combine(n int, k int) [][]int {
  path := make([]int, 0)
  ret := make([][]int, 0)
  helper(n, k, 1, path, &ret)
  return ret
}

func helper(n int, k int, start int, path []int, ret *[][]int) {
  if len(path) == k {
	tmp := make([]int, len(path))
	copy(tmp, path)
	*ret = append(*ret, tmp)
  }
  for i := start; i <= n; i++ {
	path = append(path, i)
	helper(n, k, start+i, path, ret)
	path = path[:len(path)-1]
  }
}