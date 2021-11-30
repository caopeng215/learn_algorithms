package main

import "fmt"

func main() {
  arr := []int{2,5,4,8,1,6,6,0}
  fmt.Printf("result is %v", sort(arr))
}

func sort(arr []int) []int {
  return quickSort(arr, 0, len(arr)-1)
}

func quickSort(arr []int, left int, right int) []int {
	if left < right {
      index := partition(arr, left, right)
      quickSort(arr, left, index-1)
      quickSort(arr, index+1, right)
	}
	return arr
}

func partition(arr []int, left int, right int) int {
  pivot := left
  index := left + 1
  for i:= index; i <= right; i++ {
	  if arr[pivot] > arr[i] {
	    // 交换 i和 index元素
	    arr[index], arr[i] = arr[i], arr[index]
	    index++
	  }
  }
  arr[pivot], arr[index-1] = arr[index-1], arr[pivot]
  return index - 1
}