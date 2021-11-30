package main

import "fmt"

// define struct
type MyCircularDeque struct {
  data   []int
  front  int
  rear	 int
}

// constractor MyCircularDeque
func Constructor(k int) MyCircularDeque {
  return MyCircularDeque{
	data: make([]int, k),
	front: 0,
	rear: 0,
  }
}

// 所有指针左移操作均需要加上len(this.data)，防止跨界
// Insert Front 从队列前端插入，可能跨界
func (this *MyCircularDeque) InsertFront(value int) bool {
  if this.IsFull() {
	return false
  }
  this.front = (len(this.data) + this.front - 1) % len(this.data)
  this.data[this.front] = value
  return true
}

// Insert Last 从队列后端插入
func (this *MyCircularDeque) InsertLast(value int) bool {
  if this.IsFull() {
	return false
  }
  this.data[this.rear] = value
  this.rear = (this.rear + 1) % len(this.data)
  return true
}

func (this *MyCircularDeque) DeleteFront() bool {
  if this.IsEmpty() {
	return false
  }
  this.front = (this.front + 1) % len(this.data)
  return true
}

func (this *MyCircularDeque) DeleteLast() bool {
  if this.IsEmpty() {
	return false
  }
  this.rear = (len(this.data) + this.rear -1) % len(this.data)
  return true
}

func (this *MyCircularDeque) GetFront() int {
  if this.IsEmpty() {
	return -1
  }
  return this.data[this.front]
}

func (this *MyCircularDeque) GetRear() int {
  if this.IsEmpty() {
	return -1
  }
  return this.data[(len(this.data) + this.rear - 1) % len(this.data)]
}

func (this *MyCircularDeque) IsEmpty() bool {
  if this.rear == this.front {
	return true
  }
  return false
}

func (this *MyCircularDeque) IsFull() bool {
  if (this.rear + 1) % len(this.data) == this.front {
	return true
  }
  return false
}

func main() {
	obj := Constructor(5)
	param_1 := obj.InsertFront(1)
	fmt.Printf("param_1 is: %v\n", param_1)
	param_2 := obj.InsertLast(5)
    fmt.Printf("param_2 is: %v\n", param_2)
	param_3 := obj.DeleteFront()
	fmt.Printf("param_3 is: %v\n", param_3)
	param_4 := obj.DeleteLast()
	fmt.Printf("param_4 is: %v\n", param_4)
	param_5 := obj.GetFront()
	fmt.Printf("param_5 is: %v\n", param_5)
	param_6 := obj.GetRear()
	fmt.Printf("param_6 is: %v\n", param_6)
	param_7 := obj.IsEmpty()
	fmt.Printf("param_7 is: %v\n", param_7)
	param_8 := obj.IsFull()
	fmt.Printf("param_8 is: %v\n", param_8)
}