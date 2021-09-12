package main

import "fmt"

func main() {
  obstacleGrid := [][]int{
	[]int{0,0},
	[]int{1,1},
	[]int{0,0},
  }
  res := uniquePathsWithObstacles(obstacleGrid)
  fmt.Printf("result is %v", res)
}

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
  row := len(obstacleGrid)
  col := len(obstacleGrid[0])
  dp := make([][]int, row)
  for i:=0; i< row; i++ {
	dp[i] = make([]int, col)
  }
  for i:=0; i< row && obstacleGrid[i][0] == 0; i++ {
    dp[i][0] = 1
  }
  for j:=0; j < col && obstacleGrid[0][j] == 0; j++ {
	dp[0][j] = 1
  }
  for i := 1; i < row; i++ {
	for j := 1; j < col; j++ {
	  if obstacleGrid[i][j] == 1 {
		dp[i][j] = 0
	  } else { 
		dp[i][j] = dp[i][j-1] + dp[i-1][j]
	  }
	}
  }
  return dp[row-1][col-1]
}