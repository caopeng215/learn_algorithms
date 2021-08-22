package main

import(
	"fmt"
	// "strconv"
	"sort"
)

func main() {
	nums := []int{3, 3, 0, 3}
	ret := permuteUnique(nums)
	fmt.Printf("result is: %v", ret)
}

// 方法一： 结果处优化
// func permuteUnique(nums []int) [][]int {
//     visited := map[int]bool{}
// 	path := []int{}
// 	result := [][]int{}
// 	keyhash := map[string]bool{}
//     helper(nums, path, visited, keyhash, &result)
// 	return result
// }

// func helper(nums []int, path []int, visited map[int]bool, keyhash map[string]bool, result *[][]int) {
//    if len(path) == len(nums) {
// 	 var str string 
// 	 for _, v := range path {
// 		str += strconv.Itoa(v)
// 	 }
// 	 if !keyhash[str] {
//        tmp := make([]int, len(nums))
// 	   copy(tmp, path)
// 	   *result = append(*result, tmp)
// 	   keyhash[str] = true
// 	 }
//    }
//    for i:= 0; i<len(nums); i++ {
// 	  if visited[i] {
// 		continue
// 	  }
// 	  path = append(path, nums[i])
// 	  visited[i] = true
// 	  helper(nums, path, visited, keyhash, result)
// 	  visited[i] = false
// 	  path = path[:len(path) - 1]
//    }
// }

// 方法二：
func permuteUnique(nums []int) [][]int {
    visited := map[int]bool{}
	path := []int{}
	result := [][]int{}
	// 先排序
	sort.Ints(nums)
    helper(nums, path, visited, &result)
	return result
}

func helper(nums []int, path []int, visited map[int]bool, result *[][]int) {
   if len(path) == len(nums) {
	tmp := make([]int, len(nums))
	copy(tmp, path)
	*result = append(*result, tmp)
   }
   for i:= 0; i<len(nums); i++ {
	  if visited[i] {
		continue
	  }
	  if i > 0 && nums[i] == nums[i-1] && !visited[i-1] {
		continue
	  }
	  path = append(path, nums[i])
	  visited[i] = true
	  helper(nums, path, visited, result)
	  visited[i] = false
	  path = path[:len(path) - 1]
   }
}