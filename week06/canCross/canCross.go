package main

import "fmt"

func main() {
  stones := []int{0, 1, 3, 5, 6, 8, 12, 17}
  fmt.Printf("result is %v", canCross(stones))
}

func canCross(stones []int) bool {
  m := map[int]bool{}
  return helper(m, stones, 0, 0)
}

func helper(m map[int]bool, stones []int, index int, k int) bool {
  key := index * 1000 + k
  if m[key] {
	return false
  } else {
	m[key] = true
  }
  for i:=index+1; i<len(stones); i++ {
	gap := stones[i] - stones[index]
	if gap >= k-1 && gap <= k+1 {
	  if helper(m, stones, i, gap) {
		return true
	  }
	}
	if gap > k+1 {
	  return false
	}
	if gap < k-1 {
	  continue
	}
  }
  return index == len(stones) -1
}