package week_01

type MyCircularDeque struct {
	Val   []int // data
	Count int   // 容量
}

/** Initialize your data structure here. Set the size of the deque to be k. */
func Constructor(k int) MyCircularDeque {

	return MyCircularDeque{
		Count: k,
	}
}

/** Adds an item at the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertFront(value int) bool {
	if this.IsEmpty() {
		this.Val = append(this.Val, value)
		return true

	}

	if this.IsFull() {
		return false
	}

	this.Val = append([]int{value}, this.Val...)
	return true
}

/** Adds an item at the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertLast(value int) bool {
	if this.IsEmpty() {
		this.Val = append(this.Val, value)
		return true
	}

	if this.IsFull() {
		return false
	}
	this.Val = append(this.Val, value)
	return true
}

/** Deletes an item from the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteFront() bool {
	if this.IsEmpty() {
		return false
	}

	this.Val = this.Val[1:]
	return true
}

/** Deletes an item from the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteLast() bool {
	if this.IsEmpty() {
		return false
	}

	this.Val = this.Val[0 : len(this.Val)-1]
	return true

}

/** Get the front item from the deque. */
func (this *MyCircularDeque) GetFront() int {
	if this.IsEmpty() {
		return -1
	}
	return this.Val[0]
}

/** Get the last item from the deque. */
func (this *MyCircularDeque) GetRear() int {
	if this.IsEmpty() {
		return -1
	}

	return this.Val[len(this.Val)-1]
}

/** Checks whether the circular deque is empty or not. */
func (this *MyCircularDeque) IsEmpty() bool {
	return len(this.Val) == 0
}

/** Checks whether the circular deque is full or not. */
func (this *MyCircularDeque) IsFull() bool {
	return len(this.Val) == this.Count
}

/**
 * Your MyCircularDeque object will be instantiated and called as such:
 * obj := Constructor(k)
 * param_1 := obj.InsertFront(value)
 * param_2 := obj.InsertLast(value)
 * param_3 := obj.DeleteFront()
 * param_4 := obj.DeleteLast()
 * param_5 := obj.GetFront()
 * param_6 := obj.GetRear()
 * param_7 := obj.IsEmpty()
 * param_8 := obj.IsFull()
 */

// func main() {
// 	circularDeque := Constructor(3)          // 设置容量大小为3
// 	fmt.Println(circularDeque.InsertLast(1)) // 返回 true
// 	//fmt.Println("1111", circularDeque.Len)

// 	fmt.Println(circularDeque.InsertLast(2)) // 返回 true
// 	//fmt.Println("1111", circularDeque.Len)
// 	fmt.Println(circularDeque.InsertFront(3)) // 返回 true
// 	//fmt.Println("1111", circularDeque.Len)
// 	fmt.Println(circularDeque.InsertFront(4)) // 已经满了，返回 false
// 	fmt.Println(circularDeque.GetRear())      // 返回 2
// 	fmt.Println(circularDeque.IsFull())       // 返回 true
// 	fmt.Println(circularDeque.DeleteLast())   // 返回 true
// 	fmt.Println(circularDeque.InsertFront(4)) // 返回 true
// 	fmt.Println(circularDeque.GetFront())     // 返回 4

// }
