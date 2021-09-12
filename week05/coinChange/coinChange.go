package main

import (
	"fmt"
	"math"
)

func main() {
  coins := []int{1,2,5}
  amount := 11
  fmt.Printf("result is %v", coinChange(coins, amount))
}

func coinChange(coins []int, amount int) int {
  if amount < 1 && len(coins) < 1 {
	return -1
  }
  memo := make([]int, amount+1)
  for i := 1; i <= amount; i++ {
	memo[i] = math.MaxInt32
	for _, c := range coins {
	  if i >= c && memo[i] > memo[i-c] + 1 {
		memo[i] = memo[i-c] + 1
	  }
	}
  }
  if memo[amount] == math.MaxInt32 { 
	return -1
  }
  return memo[amount]
}