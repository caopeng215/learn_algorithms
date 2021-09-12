package main

import (
	"fmt"
	"math"
)

func main() {
  k := 2
  matrix := [][]int{
	[]int{1, 0, 1},
	[]int{0, -2, 3},
  }
  fmt.Printf("result is %v:", maxSumSubmatrix(matrix, k))
}

func maxSumSubmatrix(matrix [][]int, k int) int {
  ans := math.MinInt32
  row := len(matrix)
  col := len(matrix[0])
  for left := 0; left < col; left++ {
	rSum := make([]int, row)
	for right := left; right < col; right++ {
	  for i := 0; i < row; i++ {
		rSum[i] += matrix[i][right] 
	  }
	  res := helper(rSum, k)
	  if res > ans {
		ans = res
	  }
	}
  }
  return ans
}
func helper(nums[]int, k int) int {
  max := math.MinInt32
  sum := 0
  for _, num := range nums {
	if sum > 0 {
	  sum += num
	} else {
	  sum = num
	}
	if sum > max {
	  max = sum
	}
  }
  if max <= k {
	return max
  }
  max = math.MinInt32
  for i:= 0; i< len(nums); i++ {
	sum := 0
	for j := i; j < len(nums); j++ {
	  sum += nums[j]
	  if sum > max && sum <= k {
		max = sum
	  }
	}
  }
  return max
}