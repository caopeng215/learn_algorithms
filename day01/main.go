package main

import (
	"fmt"
)

func main() {
	s1 := "abc"
	s2 := []rune(s1) // 把字符串强制转换成rune切片
	s2[0] = 'v' //  单引号
	fmt.Printf(string(s2)) //把rune切片强制转换成字符串
}