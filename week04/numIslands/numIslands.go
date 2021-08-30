package main

import "fmt"

func main() {
	grid := [][]byte{
	[]byte{"1","1","1","1","0"},
    []byte{"1","1","0","1","0"},
    []byte{"1","1","0","0","0"},
    []byte{"0","0","0","0","0"},
  }
  res := numIslands(grid)
  fmt.Printf("islands number is : %v", res)
}
func numIslands(grid [][]byte) int {
  var ret int
  return ret
}