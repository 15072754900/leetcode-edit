package main

import "fmt"

func main() {
	// 求取使用链表存储的两数相加和

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers3(l1 *ListNode, l2 *ListNode) *ListNode {
	// 包含两个链表，进行数字相加
	var num1, num2, cur uint64
	num1, num2 = 0, 0
	cur = 0
	for l1 != nil {
		if cur == 0 {
			cur++
		} else {
			cur *= 10
		}
		num1 += uint64(l1.Val) * cur
		l1 = l1.Next
	}
	cur = 0
	for l2 != nil {
		if cur == 0 {
			cur++
		} else {
			cur *= 10
		}
		num2 += uint64(l2.Val) * cur
		l2 = l2.Next
	}
	// 这里还是按照一个虚拟头结点的方法完成
	head := new(ListNode)
	ans := head
	num := num1 + num2
	fmt.Println(num1, num2, num)
	for num > 1 {
		n := num / 10
		// fmt.Println(n,num-n*10)
		cur := &ListNode{Val: int(num - n*10), Next: nil}
		// fmt.Println(cur.Val)
		ans.Next = cur
		ans = ans.Next
		num = n
	}
	res := []int{}
	for head.Next != nil {
		// fmt.Println(head.Next.Val)
		res = append(res, head.Next.Val)
		head = head.Next

	}
	return head.Next
}

// func addTwoNumbers1(l1 *ListNode, l2 *ListNode) *ListNode {
// 	var a, b int
// 	head := new(ListNode)
// 	if l1 == nil && l2 == nil {
// 		return nil
// 	} else if l1 == nil && l2 != nil {
// 		return l2
// 	} else if l2 == nil && l1 != nil {
// 		return l1
// 	}
// 	for l1.Next != nil || l2.Next != nil {
// 		cur := new(ListNode)
// 		head.Next = cur
// 		// 获取当前数据要分析是否为nil，为空则不能直接获取数据，会报错
// 		if l1 == nil && l2 != nil {
// 			a = 0
// 			b = l2.Val
// 		} else if l2 == nil && l1 != nil {
// 			a = l1.Val
// 			b = 0
// 		}
// 		if a+b < 10 {
// 			head.Val += a + b
// 		} else {
// 			higher := 0
// 			higher = a + b - (a+b)%10
// 			head.Val = (a + b) % 10
// 			head.Next.Val += higher
// 		}
// 		// 获取当前数据要分析是否为nil，为空则不能直接获取数据，会报错
// 		if l1 == nil && l2 != nil {
// 			l2 = l2.Next
// 		} else if l2 == nil && l1 != nil {
// 			l1 = l1.Next
// 		}
// 	}
// 	return head
// }

// func addTwoNumbers(l1 *ListNode, l2 *ListNode) {
// 	cur := 0
// 	higher := 0
// 	level := 0
// 	curr := 0
// 	for l1.Next != nil || l2.Next != nil {
// 		if l1.Val+l2.Val > 10 {
// 			curr = l1.Val + l2.Val - (l1.Val+l2.Val)%10
// 			higher = (l1.Val + l2.Val) % 10
// 		} else {
// 			cur = (l1.Val + l2.Val) % 10
// 		}
// 		if higher != 0 {
// 			curr += higher
// 		}
// 		higher = 0
// 		hight := 0
// 		for i := 0; i < level; i++ {
// 			hight *= 10
// 		}
// 		cur += curr * hight
// 		if l1.Next != nil {
// 			l1 = l1.Next
// 		} else {
// 			l1 = nil
// 		}
// 		if l2.Next != nil {
// 			l2 = l2.Next
// 		} else {
// 			l2 = nil
// 		}
// 	}
// 	fmt.Println(cur)
// }
