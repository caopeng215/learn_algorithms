package main

import "fmt"

func main()  {
  nums := [][]byte{
	[]byte{'1','0','1','0','0'},
	[]byte{'1','0','1','1','1'},
	[]byte{'1','1','1','1','1'},
	[]byte{'1','0','0','1','0'},
  }
  fmt.Printf("result is %v", maximalSquare(nums))
}

func maximalSquare(matrix [][]byte) int {
  m := len(matrix)
  n := len(matrix[0])
  dp := make([][]int, m+1)
  maxSlid := 0
  for i:=0; i<= m; i++ {
	dp[i] = make([]int, n+1)
  }
  for i:=1; i<=m; i++ {
	for j:=1; j<=n; j++ {
	  if matrix[i-1][j-1] == '1'{
        dp[i][j] = min(min(dp[i-1][j-1], dp[i-1][j]), dp[i][j-1]) + 1
		maxSlid = max(maxSlid, dp[i][j])
	  }
	}
  }
  fmt.Printf("dp is %v:", dp)
  return maxSlid * maxSlid
}
func max(a, b int) int {
  if a > b {
	return a
  } 
  return b
}
func min(a, b int) int {
	if a < b {
	  return a
	}
	return b
  }