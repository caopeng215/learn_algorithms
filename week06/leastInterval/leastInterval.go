package main

import "fmt"

func main() {
  tasks := []byte{'A','A','A','B','B','B'}
  n := 2
  fmt.Printf("result is %v:", leastInterval(tasks, n))
}

// func leastInterval(tasks []byte, n int) int {
//   length := len(tasks)
//   if length <= 1 {
// 	return length
//   }
//   m := map[byte]int{}
//   maxC := 0
//   for _, c := range tasks {
// 	m[c]++
// 	if maxC < m[c] {
// 	  maxC = m[c]
// 	}
//   }
//   cnt := 0
//   for _, v := range m {
// 	if v == maxC {
// 	  cnt++
// 	}
//   }
//   return max(length, (maxC - 1) * (n+1) + cnt)
// }

func leastInterval(tasks []byte, n int) int {
  length := len(tasks)
  if length <= 1 {
	return length
  }
  m := make([]int, 26)
  maxC := 0
  for _, c := range tasks {
	m[c-'A']++
  }
  for _, v := range m {
	if v > maxC {
	  maxC = v
	}
  }
  cnt := 0
  for _, v := range m {
	if v == maxC {
	  cnt++
	}
  }
  return max(length, (maxC - 1) * (n+1) + cnt)
}
func max(a, b int) int {
  if a > b {
	return a
  }
  return b
}