func addTwoNumbers(l1 *ListNode, l2 *ListNode) (out *ListNode) {
    head := &ListNode{}
    focus := head
    carry := 0

	for l1 != nil || l2 != nil || carry != 0 {
		l1Val := 0
		l2Val := 0
		if l1 != nil {
			l1Val = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			l2Val = l2.Val
			l2 = l2.Next
		}
		currentSum := l1Val + l2Val + carry
		currentVal := currentSum % 10
		carry = currentSum / 10

        current := ListNode{Val: currentVal}
        focus.Next = &current

        focus = focus.Next
	}
    return head.Next
}
