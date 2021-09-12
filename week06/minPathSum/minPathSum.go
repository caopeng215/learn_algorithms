package main

import "fmt"

func main() {
  grid := [][]int{ 
    []int{1, 2, 3},
	[]int{4, 5, 6},
  }
  ans := minPathSum(grid)
  fmt.Printf("ans is %v:", ans)
}

// 方法一
// func minPathSum(grid [][]int) int {
//   m := len(grid)
//   n := len(grid[0])
//   dp := make([][]int, m)
//   for i:= 0; i<m; i++ {
// 	dp[i] = make([]int, n)
//   }
//   for i:=0; i<m; i++ {
// 	for j:=0; j<n; j++ {
// 	  if i == 0 && j == 0 {
// 		dp[i][j] = grid[i][j]
// 	  } else if i == 0 {
// 		dp[i][j] = dp[i][j-1] + grid[i][j]
// 	  } else if j == 0 {
// 		dp[i][j] = dp[i-1][j] + grid[i][j]
// 	  } else {
// 	    dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
// 	  }
// 	}
//   }
//   fmt.Printf("dp is %v:", dp)
//   return dp[m-1][n-1]
// }

// 方法二
func minPathSum(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	for i:=0; i<m; i++ {
	  for j:=0; j<n; j++ {
		if i == 0 && j == 0 {
		  continue
		} else if i == 0 {
			grid[i][j] = grid[i][j-1] + grid[i][j]
		} else if j == 0 {
			grid[i][j] = grid[i-1][j] + grid[i][j]
		} else {
			grid[i][j] = min(grid[i-1][j], grid[i][j-1]) + grid[i][j]
		}
	  }
	}
	return grid[m-1][n-1]
  }
func min(a, b int)int {
  if a < b {
	return a
  }
  return b
}