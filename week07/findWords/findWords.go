package main

import "fmt"

type TrieNode struct {
  word string
  children  [26]*TrieNode
}

func main() {
	board := [][]byte{
	  []byte{'o','a','a','n'},
	  []byte{'e','t','a','e'},
	  []byte{'i','h','k','r'},
	  []byte{'i','f','l','v'},
	}
	words := []string{"oath","pea","eat","rain"}
	fmt.Print("result is: %v", findWords(board, words))
}

func findWords(board [][]byte, words []string) []string {
  root := &TrieNode{}
  for _, word := range words {
	node := root
	for _, c := range word {
	  if node.children[c-'a'] == nil {
		node.children[c-'a'] = &TrieNode{}
	  }
	  node = node.children[c-'a'] 
	}
	node.word = word
  }
  result := make([]string, 0)
  for i := 0; i < len(board); i++ {
	for j := 0; j < len(board[0]); j++ {
		dfs(i, j, board, root, &result)
	}
  }
  return result
}
func dfs(row int, col int, board [][]byte, node *TrieNode, result *[]string) {
  if row < 0 || col < 0 || row == len(board) || col == len(board[0]) {
	return
  }
  c := board[row][col]
  if c == '#' || node.children[c - 'a'] == nil {
	return
  }
  node = node.children[c - 'a']
  if node.word != "" {
    *result = append(*result, node.word)
	node.word = ""
  }
  board[row][col] = '#'
  dfs(row+1, col, board, node, result)
  dfs(row, col-1, board, node, result)
  dfs(row, col+1, board, node, result)
  dfs(row-1, col, board, node, result)
  board[row][col] = c
} 