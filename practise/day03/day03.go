package main

import "fmt"

// 链表结构体
type ListNode struct{
  Val int
  Next *ListNode
}

func main() {
  // 新建链表
  node4 := ListNode{Val: 4, Next: nil}
  node3 := ListNode{Val: 3, Next: &node4}
  node2 := ListNode{Val: 2, Next: &node3}
  head := &ListNode{Val: 1, Next:&node2}
  // 遍历链表
  // node := head.Next
  // for node != nil {
  //   fmt.Printf("list is %v:", node.Val)
  //   node = node.Next
  // }
  // 两两交换链表数字
  ans := swapPairs(head)
  // 打印swap后的链表
  node := ans
  for node != nil {
    fmt.Printf("list is %v:", node.Val)
    node = node.Next
  }
}

// 方法一：交换指针
// swap pairs
// func swapPairs(head *ListNode) *ListNode {
//   virtualNode := &ListNode{Val: 0, Next: head}
//   tempNode := virtualNode
//   for tempNode.Next != nil && tempNode.Next.Next != nil {
// 	  pre := tempNode.Next
// 	  after := tempNode.Next.Next
// 	  tempNode.Next = after
// 	  pre.Next = after.Next
// 	  after.Next = pre
// 	  tempNode = pre
//   }
//   return virtualNode.Next
// }

// 方法二
// head(1) --> head.Next(2)-->(3)-->(4)
// virNode = head.Next(2)
// head(1) --> swapPairs((3)-->(4))
// virNode.Next --> head(参数)
func swapPairs(head *ListNode) *ListNode {
  if head == nil || head.Next == nil {
	return head
  }
  virNode := head.Next
  head.Next = swapPairs(virNode.Next)
  virNode.Next = head
  return virNode
}