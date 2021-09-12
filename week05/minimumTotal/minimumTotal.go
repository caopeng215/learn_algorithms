package main

import "fmt"

func main() {
  triangle := [][]int{
	[]int{2},
	[]int{3,4},
	[]int{6,5,7},
	[]int{4,1,8,3},
  }
  ans := minimumTotal(triangle)
  fmt.Printf("ans is %v:", ans)
}
// 方法一
// var memo [][]int
// func minimumTotal(triangle [][]int) int {
//   m := len(triangle)
//   n := len(triangle[len(triangle)-1])
//   memo = make([][]int, m)
//   for i:=0; i<m; i++ {
// 	memo[i] = make([]int, n)
//   }
//   return dfs(triangle, 0, 0)
// }

// func dfs(triangle[][]int, row int, col int) int {
//   if len(triangle) == row {
// 	return 0
//   }
//   if memo[row][col] != 0 {
//     return memo[row][col]
//   }
//   memo[row][col] = min(dfs(triangle, row+1, col), dfs(triangle, row+1, col+1)) + triangle[row][col]
//   return memo[row][col]
// }
func min(a, b int)int {
  if a > b {
	return b
  }
  return a
}

// 方法二
func minimumTotal(triangle [][]int) int {
  m := len(triangle)
  dp := make([][]int, m+1)
  for i := 0; i <= m; i++ {
	dp[i] = make([]int, m+1)
  }
  for i:= m-1; i >=0; i-- {
	for j := 0; j<=i; j++ {
	  dp[i][j] = min(dp[i+1][j],dp[i+1][j+1])+triangle[i][j]
	}
  }
  return dp[0][0]
}