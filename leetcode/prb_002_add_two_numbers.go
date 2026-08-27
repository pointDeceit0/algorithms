package leetcode


// Definition for singly-linked list.
type ListNode struct {
    Val int
    Next *ListNode
}


func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	next := l1
    beginning := next
	remain := 0
	prev := l1

	for l1 != nil && l2 != nil {
		prev = next
		tmp := l1.Val + l2.Val + remain
		next.Val = tmp % 10
		remain = tmp / 10
		next = next.Next
		l1, l2 = l1.Next, l2.Next
	}

	for l1 != nil {
		prev = next
		tmp := l1.Val + remain
		next.Val = tmp % 10
		remain = tmp / 10
		next = next.Next
		l1 = l1.Next
	}

	if l2 != nil {
		prev.Next = l2
		next = prev.Next
	}
	for l2 != nil {
		prev = next
		tmp := l2.Val + remain
		next.Val = tmp % 10
		remain = tmp / 10
		next = next.Next
		l2 = l2.Next
	}

	if remain != 0 {
		prev.Next = &ListNode{Val: remain, Next: nil}
		next = prev.Next
	}

	return beginning
}
