package main
import "fmt"

func main() {
  a := []int{1,2,3,2,1}
  b := []int{3,2,1,4,7}
  fmt.Printf("ans is: %v", findLength(a, b))
}

// func findLength(nums1 []int, nums2 []int) int {
//   ans := 0
//   m := len(nums1)
//   n := len(nums2)
//   dp := make([][]int, m+1)
//   for i:=0; i <= m; i++ {
// 	dp[i] = make([]int, n+1)
//   }
  
//   for i:=1; i<=m; i++ {
// 	for j:=1; j<=n;j++ {
// 	  if nums1[i-1] == nums2[j-1] {
// 		dp[i][j] = dp[i-1][j-1] + 1
// 	  } else {
// 		dp[i][j] = 0
// 	  }
// 	  if dp[i][j] > ans {
// 		ans = dp[i][j]
// 	  }
// 	}
//   }
//   return ans
// }

func findLength(nums1 []int, nums2 []int) int {
	ans := 0
	m := len(nums1)
	n := len(nums2)
	dp := make([]int, n+1)
	
	for i:=1; i<=m; i++ {
	  for j:=n; j>0; j-- {
		if nums1[i-1] == nums2[j-1] {
		  dp[j] = dp[j-1] + 1
		} else {
		  dp[j] = 0
		}
		if dp[j] > ans {
		  ans = dp[j]
		}
	  }
	}
	return ans
  }