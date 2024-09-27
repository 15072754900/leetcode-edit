package main

import "fmt"

func main() {
	// 求取使用链表存储的两数相加和

}

type ListNode struct {
	val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) {
	cur := 0
	higher := 0
	level := 0
	curr := 0
	for l1.Next != nil || l2.Next != nil {
		if l1.val+l2.val > 10 {
			curr = l1.val + l2.val - (l1.val+l2.val)%10
			higher = (l1.val + l2.val) % 10
		} else {
			cur = (l1.val + l2.val) % 10
		}
		if higher != 0 {
			curr += higher
		}
		higher = 0
		hight := 0
		for i := 0; i < level; i++ {
			hight *= 10
		}
		cur += curr * hight
		if l1.Next != nil {
			l1 = l1.Next
		} else {
			l1 = nil
		}
		if l2.Next != nil {
			l2 = l2.Next
		} else {
			l2 = nil
		}
	}
	fmt.Println(cur)
}
