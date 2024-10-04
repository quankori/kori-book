package sumtwonumber

type ListNode struct {
	Value int
	Next  *ListNode
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	current := dummy
	carry := 0

	for l1 != nil || l2 != nil {
		var x, y int
		if l1 != nil {
			x = l1.Value
			l1 = l1.Next
		}
		if l2 != nil {
			y = l2.Value
			l2 = l2.Next
		}

		sum := carry + x + y
		carry = sum / 10
		current.Next = &ListNode{Value: sum % 10}
		current = current.Next
	}

	if carry > 0 {
		current.Next = &ListNode{Value: carry}
	}

	return dummy.Next
}

func createList(numbers []int) *ListNode {
	if len(numbers) == 0 {
		return nil
	}
	head := &ListNode{Value: numbers[0]}
	current := head
	for _, v := range numbers[1:] {
		current.Next = &ListNode{Value: v}
		current = current.Next
	}
	return head
}
