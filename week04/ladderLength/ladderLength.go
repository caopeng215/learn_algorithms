package main

import "fmt"

func main() {
  begin := "hit"
  endWord := "cog"
  wordList := []string{"hot", "dot", "dog", "lot", "log", "cog"}
  ret := ladderLength(begin, endWord, wordList)
  fmt.Printf("result is %v", ret)
}

func ladderLength(beginWord string, endWord string, wordList []string) int {
  wordMap := map[string]bool{}
  for _, val := range wordList {
	wordMap[val] = true
  }
  if _, ok := wordMap[endWord]; ok && !wordMap[endWord] {
	return -1
  }
  queue := []string{beginWord}
  level := 0
  for len(queue) > 0 {
	size := len(queue)
	for i := 0; i < size; i++ {
		word := queue[0]
		queue = queue[1:]
		for j := 0; j < len(word); j++ {
		  for k := 'a'; k <= 'z'; k++ {
			newWord := word[:j] + string(k) + word[j+1:]
			if wordMap[newWord] {
			  wordMap[newWord] = false
			  queue = append(queue, newWord)
			}
		  }
		}
	}	
	level++
  }
  return level
}