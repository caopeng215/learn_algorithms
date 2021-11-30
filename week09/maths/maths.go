package main

import (
	"fmt"
	"strconv"
)

func main() {
  a := 1048575.88
  fmt.Print(a/1024)
  origin, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", a/1024), 64)
  fmt.Printf("origin: %v", origin)
}