package removeNthFromEnd

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// 思路：先获取链表的全部数量，然后建一个虚拟头结点，然后建立一个链表进行处理
	cur := head
	total := 0
	for cur != nil {
		total++
		cur = cur.Next
	}
	n = total - n + 1
	dummy := &ListNode{Next: head}
	now := dummy
	for n > 0 {
		now = now.Next
		n--
	}
	now.Next = now.Next.Next
	return dummy.Next
}

func removeNthFromEnd1(head *ListNode, n int) *ListNode {
	// 使用动态双指针方法进行处理，首先，需要进行n次将链表走n次，然后两个链表同时走，直至第一个到终点，删除第二个链表后面的内容
	dummy := &ListNode{Next: head}
	first, second := dummy, dummy

	for i := 0; i <= n; i++ {
		first = first.Next
	}
	for first != nil {
		first = first.Next
		second = second.Next
	}
	second.Next = second.Next.Next

	return dummy.Next
}
