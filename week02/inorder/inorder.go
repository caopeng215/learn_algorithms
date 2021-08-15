package main

import (
	"fmt"
	// "math/rand"
)

type TreeNode struct{
	Val int
	Left *TreeNode
	Right *TreeNode
}

func NewBuildTree(i int, nums []int) *TreeNode {
  node := &TreeNode{Val: nums[i]}
  if i < len(nums) && 2 * i + 1 < len(nums) {
	  node.Left = NewBuildTree(2 * i + 1, nums)
  }
  if i < len(nums) && 2 * i + 2 < len(nums) {
	  node.Right = NewBuildTree(2 * i + 2, nums)
  }
  return node
}

// 中序遍历
func inorder(root *TreeNode, res *[]int) {
	if root == nil {
	  return
	}
	inorder(root.Left, res)
	*res = append(*res, root.Val)
	inorder(root.Right, res)
}

var i int = -1
func BuildTree(nums []int) *TreeNode {
	i++
	var node *TreeNode
	if i >= len(nums) {
	  return nil
	} else {
	  node = new(TreeNode)
	  node.Val = nums[i]
	  node.Left = BuildTree(nums)
	  node.Right = BuildTree(nums)
	}
	return node
}

func main() {
	// demo [1,4,2,3,8,6,5] --> [3,4,8,1,6,2,5]
	res := []int{}

	// 使用随机数生成数组
    // nums := make([]int, 10)
    // for i:= 0; i < 10; i++ {
    //   nums[i] = rand.Intn(20) +1 // rand.Intn(20) 用来生成随机数，区间0~20
    // }

	nums := []int{ 1,4,2,3,8,6,5 }

	// 使用BuildTree方法构建二叉树
	// root := BuildTree(nums)

	// 使用NewBuildTree方法构建二叉树(验证正确)
	root := NewBuildTree(0, nums)

	// 递归中序遍历
	inorder(root, &res)
	fmt.Printf("tree:%v:", res)
 }