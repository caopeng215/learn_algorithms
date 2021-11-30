package main

import "fmt"

type ListNode struct{
  Val  int
  Next *ListNode
}

func main() {
  // 构建链表1
  // 手动构建
  // l3 := ListNode{ Val: 4, Next: nil }
  // l2 := ListNode{ Val: 2, Next: &l3 }
  // l1 := &ListNode{ Val: 1, Next: &l2 }
  l1 := constractorListNode([]int{1, 2, 4})
  // 构建链表2
  // 手动构建
  // l5 := ListNode{ Val: 5, Next: nil }
  // l4 := ListNode{ Val: 3, Next: &l5 }
  // l := &ListNode{ Val: 0, Next: &l4 }
  // 函数构建
  l2 := constractorListNode([]int{0, 3, 5})
  printListNode(mergeTwoList(l1, l2))
}

func constractorListNode(arrs []int) *ListNode {
  list := make([]ListNode, len(arrs))
  for i:= len(arrs) - 1; i >= 0; i-- {
    if i == len(arrs) -1 {
	  list[i] = ListNode{Val: arrs[i], Next: nil}
	} else {
	  list[i] = ListNode{Val: arrs[i], Next: &list[i+1]}
	}
  }
  // head = list[0]
  return &list[0]
}

// print List
func printListNode(head *ListNode) {
  fmt.Println("begin print list node...")
  for head != nil {
	fmt.Printf("node is: %v\n", head.Val)
	head = head.Next
  }
}

// 合并2个有序链表
// 方法一
// 采用递归方法
func mergeTwoList(l1 *ListNode, l2 *ListNode) *ListNode {
  if l1 == nil {
	return l2
  }
  if l2 == nil {
	return l1
  }
  if l1.Val < l2.Val {
    l1.Next = mergeTwoList(l1.Next, l2)
	return l1
  } else {
	l2.Next = mergeTwoList(l1, l2.Next)
    return l2
  }
}

// 方法二：
// 辅助头节点实现
// func mergeTwoList(l1 *ListNode, l2 *ListNode) *ListNode {
//   virNode := &ListNode{}
//   node := virNode
//   for l1 != nil && l2 != nil {
// 	  if l1.Val < l2.Val {
// 		node.Next = l1
// 		l1 = l1.Next
// 	  } else {
// 		node.Next = l2
// 		l2 = l2.Next
// 	  }
// 	  node = node.Next
//   }
//   if l1 == nil {
// 	node.Next = l2
//   }
//   if l2 == nil {
// 	node.Next = l1
//   }
//   return virNode.Next
// }