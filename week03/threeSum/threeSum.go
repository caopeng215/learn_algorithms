package main

import(
	"fmt"
	"sort"
)

func main() {
  nums := []int{-1,0,1,2,-1,-4} // -4, -1, -1, 0, 1, 2
  ret := threeSum(nums)
  fmt.Printf("value is:%v", ret)
}

func threeSum(nums []int) [][]int {
	n := len(nums)
	result := [][]int{}
	if len(nums) < 3 {
		return result
	}
	sort.Ints(nums)
    for i := 0; i < n; i++ {
	  k := n-1
	  if i > 0 && nums[i] == nums[i-1] {
		continue
	  }
	  for j:=i+1; j<n; j++ {
        if j > i+1 && nums[j] == nums[j-1] {
		  continue
		}
		for j < k && nums[i] + nums[j] + nums[k] > 0 {
		  k--
		}
		if j == k {
		  break
		}
		if j < k && nums[i] + nums[j] + nums[k] == 0 {
		  result = append(result, []int{ nums[i], nums[j], nums[k] })
		}
	  }
	}
	return result
  }