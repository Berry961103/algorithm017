package week_01

//  [1,8,6,2,5,4,8,3,7]
func maxArea(height []int) int {
	max := 0                 // 记录最大的面积
	i, j := 0, len(height)-1 // i 左边的棒子位置，j右边的棒子位置
	for i != j {             // 一直比较到  i 和j是同一个
		left, right := height[i], height[j]
		sMax := (j - i) * minHeight(left, right) //  （j-i）j到i的距离，也叫底，最小的高度，算出面积

		if sMax > max {
			max = sMax
		}

		// 移动棒子
		if left > right { // 移动小的那边棒子的位置
			j--
		} else {
			i++
		}
	}
	return max
}
func minHeight(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

/**
 * Definition for singly-linked list.

 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func ReverseList(head *ListNode) *ListNode {
	var newHead, next *ListNode
	for head != nil {
		next = head.Next    // next =2,3  3  nil
		head.Next = newHead // head=1   2,1  3,2,1
		newHead = head      // new 1    2,3 3,2,1
		head = next         // head 2,3  3 nil
	}
	return newHead
}
func ReverseList2(head *ListNode) *ListNode {
	if nil == head {
		return head
	}
	cur, next := head, head.Next
	for nil != next {
		cur, next, next.Next = next, next.Next, cur
		// cur = next
		// next = next.Next
		// next.Next = cur
	}
	head.Next = nil
	return cur
}
func ReverseList3(head *ListNode) *ListNode {
	if nil == head || head.Next == nil {
		return head
	}

	cur := ReverseList3(head.Next)
	head.Next.Next = head
	head.Next = nil
	return cur
}
