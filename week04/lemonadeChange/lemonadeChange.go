package main

import "fmt"

func main() {
  bills := []int{5,5,10}
  ret := lemonadeChange(bills)
  fmt.Printf("result is: %v", ret)
}

// func lemonadeChange(bills []int) bool {
//   five, ten := 0, 0
//   for _, bill := range bills {
// 	if bill == 5 {
// 	  five++
// 	} else if bill == 10 {
// 	  if five < 0 {
// 	    return false
// 	  }
// 	  five--
// 	  ten++
// 	} else {
// 	  if five > 0 && ten > 0 {
// 	    five--
// 	    ten--
// 	  } else if five > 3 {
// 	    five -= 3
// 	  } else {
// 	    return false
// 	  }
// 	}
//   }
//   return true
// }

func lemonadeChange(bills []int) bool {
	billMap := map[int]int{}
	for _, bill := range bills {
	  if bill == 5 {
		billMap[5]++
	  } else if bill == 10 {
		if billMap[5] < 0 {
		  return false
		}
		billMap[5]--
		billMap[10]++
	  } else {
		if billMap[5] > 0 && billMap[10] > 0 {
		  billMap[5]--
		  billMap[10]--
		} else if billMap[5] > 3 {
		  billMap[5] -= 3
		} else {
		  return false
		}
	  }
	}
	return true
  }