package medium


// Definition for singly-linked list.
type ListNode struct {
    Val int
    Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    var next *ListNode

	sum := l1.Val + l2.Val
	current := sum % 10
	addOne := sum / 10

	if addOne == 1 {
		if l1.Next == nil {
			l1.Next = createList(0, nil)
		}
		l1.Next.Val += 1
	}

	if l1.Next != nil && l2.Next != nil {
		next = addTwoNumbers(l1.Next, l2.Next)
	} else if l1.Next != nil {
		next = addTwoNumbers(l1.Next, createList(0, nil))
	} else if l2.Next != nil {
		next = addTwoNumbers(createList(0, nil), l2.Next)
	}

	return createList(current, next)
}

func createList(val int, next *ListNode) *ListNode {
	return &ListNode{
		Val: val,
		Next: next,
	}
}