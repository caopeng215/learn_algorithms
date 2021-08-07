type MyCircularDeque struct {
   data []int
   front int
   rear int
}


/** Initialize your data structure here. Set the size of the deque to be k. */
func Constructor(k int) MyCircularDeque {
  return MyCircularDeque{
      data: make([]int, k+1, k+1),
      front: 0,
      rear: 0,
  }
}


/** Adds an item at the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertFront(value int) bool {
  if this.IsFull() {
      return false
  }
  this.front = (len(this.data) + this.front - 1) % len(this.data)
  this.data[this.front] = value
  return true
}


/** Adds an item at the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertLast(value int) bool {
  if this.IsFull() {
      return false
  }
  this.data[this.rear] = value
  this.rear = (this.rear + 1) % len(this.data)
  return true
}


/** Deletes an item from the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteFront() bool {
  if this.IsEmpty() {
      return false
  }
  this.front = (this.front + 1) % len(this.data)
  return true
}


/** Deletes an item from the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteLast() bool {
  if this.IsEmpty() {
      return false
  }
  this.rear = (len(this.data) + this.rear - 1) % len(this.data)
  return true
}


/** Get the front item from the deque. */
func (this *MyCircularDeque) GetFront() int {
  if this.IsEmpty() {
    return -1
  }
  return this.data[this.front]
}


/** Get the last item from the deque. */
func (this *MyCircularDeque) GetRear() int {
  if this.IsEmpty() {
      return -1
  }
  return this.data[(len(this.data) + this.rear - 1) % len(this.data)]
}


/** Checks whether the circular deque is empty or not. */
func (this *MyCircularDeque) IsEmpty() bool {
  return this.rear == this.front
}


/** Checks whether the circular deque is full or not. */
func (this *MyCircularDeque) IsFull() bool {
  return this.front == (this.rear + 1) % len(this.data)
}
