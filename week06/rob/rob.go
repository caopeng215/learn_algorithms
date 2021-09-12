package main

import "fmt"

func main()  {
  nums := []int{1, 2, 3, 1}
  fmt.Printf("result is %v:", rob(nums))
}

// 方法一： dp动态规划
// func rob(nums []int) int {
//   if len(nums) == 1 {
// 	return nums[0]
//   }
//   nums1 := nums[1:]
//   nums2 := nums[:len(nums)-1]
//   dp := make([]int, len(nums))
//   dp[1] = nums1[0]
//   for i := 2; i < len(nums); i++ {
// 	dp[i] = max(dp[i-1], dp[i-2]+nums1[i-1])
//   }
//   cnt := dp[len(nums) - 1]
//   dp1 := make([]int, len(nums))
//   dp1[1] = nums2[0]
//   for i := 2; i < len(nums); i++ {
// 	dp1[i] = max(dp1[i-1], dp1[i-2]+nums2[i-1])
//   }
//   if dp1[len(nums) - 1] > cnt {
// 	cnt = dp1[len(nums) - 1]
//   }
//   return cnt
// }
func rob(nums []int) int {
	if len(nums) == 0 {
	  return 0
	}
	if len(nums) == 1 {
	  return nums[0]
	}
    return max(maxRob(nums[1:]), maxRob(nums[:len(nums)-1]))
  }
func maxRob(nums[]int) int {
  pre, cur := 0, 0
  for _, v := range nums {
	pre, cur = cur, max(pre + v, cur)
  }
  return cur
}
func max(a, b int) int {
  if a > b {
	return a
  }
  return b
}