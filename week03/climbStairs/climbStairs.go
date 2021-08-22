package main

import (
    "fmt"
)

func main() {
  n := 10
  ret := climbStairs(n)
  fmt.Printf("result is: %v", ret)
}

func climbStairs(n int) int {
  // 方法一：递归
  // if n == 1 || n == 2 {
  //   return n
  // }
  // return climbStairs(n-1) +  climbStairs(n-2)
  
  // 方法二: 迭代
  // sum := make([]int, n+1)
  // for i := 1; i < n+1; i++ {
  //   if i == 1 || i == 2 {
  //     sum[i] = i
  //   } else {
  //     sum[i] = sum[i-1] + sum[i-2]
  //   }
  // }
  // return sum[n]

  // 方法三：优化迭代
  p, q, r := 0, 0, 1
  for i := 0; i < n; i++ {
    p = q
    q = r
    r = p + q
  }
  return r
}