package main

import "fmt"

func main() {
  n := 2
  ret := generateParenthesis(n)
  fmt.Printf("result is:%v", ret)
}

func generateParenthesis(n int) []string {
  res := []string{}
  helper("", 0, 0, n, &res)
  return res
}

func helper(str string, left int, right int, n int, res *[]string) {
  if left == n && right == n {
	*res = append(*res, str)
	return
  }
  if left < right {
	return
  }
  if left < n {
	helper(str+"(", left+1, right, n, res)
  }
  if right < left {
	helper(str+")", left, right+1, n, res)
  }
}