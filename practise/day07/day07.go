package main

import (
	"fmt"
	"sort"
)

func main() {
  nums1 := []int{1,2,2,1}
  nums2 := []int{2,2}
  fmt.Printf("ans is: %v", intersect(nums1, nums2))
}

func intersect(nums1 []int, nums2 []int) []int {
  p := 0
  q := 0
  insections := []int{}
  sort.Ints(nums1)
  sort.Ints(nums2)
  for p < len(nums1) && q < len(nums2) {
    if nums1[p] < nums2[q] {
	  p++
	} else if nums2[q] < nums1[p] {
	  q++
	} else {
	  insections = append(insections, nums1[p])
	  p++
	  q++
	}
  }
  return insections
}