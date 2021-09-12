package main

import (
	"fmt"
)

func main() {
  commands := []int{4, -1, 4, -2, 4}
  obstacles := [][]int{[]int{2, 4}}
  ret := robotSim(commands, obstacles)
  fmt.Printf("result is %v", ret)
}

func robotSim(commands []int, obstacles [][]int) int {
  res := 0
  rotateX := []int{0, 1, 0, -1}
  rotateY := []int{1, 0, -1, 0}
  curPos := 0
  curX, curY := 0, 0
  obsMap := map[[2]int]bool{}
  for _, val := range obstacles {
	obsMap[[2]int{val[0], val[1]}] = true
  }
  for i := 0; i < len(commands); i++ {
	if commands[i] == -1 {
	  curPos = (curPos + 1) % 4
	} else if commands[i] == -2 {
	  curPos = (curPos + 3) % 4
	} else {
	  for j:=1; j <= commands[i]; j++ {
        nx := curX + rotateX[curPos]
	    ny := curY + rotateY[curPos]
		if obsMap[[2]int{nx, ny}] {
		  break
		}
		curX = nx
		curY = ny
		if curX * curX + curY * curY > res {
		  res = curX * curX + curY * curY
		}
	  }
	}
  }
  return res
}