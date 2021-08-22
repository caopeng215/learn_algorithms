package main

import(
	"fmt"
)

func main()  {
	n := 3
	ret := make([]string, 0)
	ret = generateParenthesis(n)
	fmt.Printf("result:", ret)
}

func generateParenthesis(n int) []string {
  result := make([]string, 0)
  deepFind("", 0, 0, n, &result)
  return result
}

func deepFind(str string, p int, q int, n int, ret *[]string) {
  if p == n && q == n {
	  *ret = append(*ret, str)
	  return
  }
  if p < q {
	  return
  }
  if p < n {
	  deepFind(str + "(", p+1, q, n, ret)
  }
  if q < p {
	  deepFind(str + ")", p, q+1, n, ret)
  }
}