package main

import "fmt"

func main() {
  nums := []int{1,2,3,4,5,6,7}
  k := 10
  arr := twoSum(nums, k)
  fmt.Printf("arr is %d:", arr)
}

// 优化代码 
// 只需要for遍历一次
func twoSum(list []int, target int) []int {
    mm := make(map[int]int, len(list))
	for i, v := range list {
      if j, ok := mm[target-v]; ok {
        return []int{j, i}  
	  } else {
		mm[v] = i
	  }
    }
	return []int{}
}