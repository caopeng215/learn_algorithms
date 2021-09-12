package main

import "fmt"

func main() {
  s := "aaa"
  fmt.Printf("result is %v:", countSubstrings(s))
}

// func countSubstrings(s string) int {
// 	length := len(s)
// 	dp := make([][]bool, length)
// 	for i:=0; i< length; i++ {
// 	  dp[i] = make([]bool, length)
// 	}
// 	cnt := 0
// 	for j:=0; j<length; j++ {
// 	  for i:=0; i<=j; i++ {
// 		if s[i] == s[j] && (j-i<2 || dp[i+1][j-1]) {
// 		  dp[i][j] = true
// 		  cnt++
// 		}
// 	  }
// 	}
// 	return cnt
//   }

// 方法二： 中心扩散法
  func countSubstrings(s string) int {
    length := len(s) * 2 - 1
	cnt := 0
	for center:=0; center < length; center++ {
	  left := center / 2
	  right	:= left + center % 2
	  for left >= 0 && right < len(s) && s[left] == s[right] {
		cnt++
		left--
		right++
	  }
	}
	return cnt
  }