package main

import "fmt"

func main() {
  m := 2
  nums := []int{7,2,5,10,8}
  fmt.Printf("result is %v", splitArray(nums, m))
}

func splitArray(nums []int, m int) int {
  // m 确定子数组个数
  // 方法：二分查找，子数组和的边界[max(nums), sums(nums)]]
  // 通过二分查找找到mid值，即刚好实现数组和最少
  lmax := nums[0]
  rmax := 0
  for _, num := range nums {
	if lmax < num {
	  lmax = num
	}
	rmax += num
  }
  for lmax < rmax {
	cnt := 1
	tmp := 0
	mid := lmax + (rmax - lmax) / 2
	for _, num := range nums {
      tmp += num
	  // tmp > mid 说明已找到一个子数组边界
	  if tmp > mid {
		tmp = num
		cnt++
	  }
	}
	// 子数组个数大于给定的m值
	if cnt > m {
	  lmax = mid + 1
	} else {
	  rmax = mid
	}
  }
  return lmax
}