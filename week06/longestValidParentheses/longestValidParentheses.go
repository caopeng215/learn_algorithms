package main

import "fmt"

func main() {
  s := ")()())"
  fmt.Printf("result is %v", longestValidParentheses(s))
}

func longestValidParentheses(s string) int {
  length := len(s)
  dp := make([]int, length)
  maxCnt := 0
  for i:=0;i<length;i++{
	if s[i] == ')' {
	  if i-1 >= 0 && s[i-1] == '(' {
		dp[i] = 2
        if i - 2 >= 0 {
		  dp[i] = dp[i-2] + dp[i]
		}
	  }
	  if i-1 >= 0 && dp[i-1] > 0 {
		if i-dp[i-1]-1 >= 0 && s[i-dp[i-1]-1] == '(' {
		  dp[i] = dp[i-1] + 2
		  if i-dp[i-1]-2 >= 0 {
		    dp[i] = dp[i] + dp[i-dp[i-1]-2]
		  }
		}
	  }
	}
  }
  for _, v := range dp {
	if v > maxCnt {
	  maxCnt = v
	}
  }
  return maxCnt
}
