package main

import "fmt"

func main() {
  word1 := "horse"
  word2 := "ros"
  fmt.Printf("result is %v", minDistance(word1, word2))
}

func minDistance(word1 string, word2 string) int {
  m := len(word1)
  n := len(word2)
  dp := make([][]int, m+1)
  for i := 0; i <= m; i++ {
	dp[i] = make([]int, n+1)
	dp[i][0] = i
  }
  for i:= 0; i<= n; i++ {
	dp[0][i] = i
  }
  for i := 1; i <=m; i++ {
	for j := 1; j <= n; j++ {
	  if word1[i-1] == word2[j-1] {
		dp[i][j] = dp[i-1][j-1]
	  } else {
	    dp[i][j] = min(min(dp[i-1][j-1],dp[i][j-1]), dp[i-1][j]) + 1
	  }
	}
  }
  return dp[m][n]
}

func min(a, b int) int {
  if a > b {
	return b
  }
  return a
}
