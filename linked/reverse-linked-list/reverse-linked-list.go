package linked

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var pre *ListNode
	cur := head
	for cur != nil {
		// pre, cur, cur.Next = cur, cur.Next, pre
		next := head.Next
		head.Next = pre
		pre = head
		head = next
	}
	return pre
}
