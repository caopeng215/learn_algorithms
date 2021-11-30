package main

import "fmt"

func main() {
  str := "loveleetcode"
  fmt.Printf("ans is: %v", firstUniqChar(str))
}

// func firstUniqChar(s string) int {
//   m := make(map[rune]int)
//   for _, c := range s {
// 	m[c]++
//   }
//   for i, c := range s {
// 	if m[c] == 1 {
// 	  return i
// 	}
//   }
//   return -1
// }

func firstUniqChar(s string) int {
	index := [26]int{}
	for i := 0; i < 26; i++ {
	  index[i] = -1
	}
	for i := 0; i < len(s); i++ {
	  if index[s[i]-'a'] == -1 {
		index[s[i]-'a'] = i
	  } else {
		index[s[i]-'a'] = -2
	  }
	}
	for i:= 0; i<len(index); i++ {
	  if index[i] > 0 {
	    if index[i] > ans {
		  
		}
	  }
	}
	return -1
  }