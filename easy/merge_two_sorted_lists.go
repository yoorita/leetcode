package main

/*Definition for singly-linked list.*/
type ListNode struct {
	Val int
	Next *ListNode
}

 func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    if list2 != nil && list1 != nil {
		if list1.Val == list2.Val {
			return &ListNode{
				Val: list1.Val,
				Next: &ListNode{
					Val: list2.Val,
					Next: mergeTwoLists(list1.Next, list2.Next),
				},
			}
		} else if list1.Val < list2.Val {
			return &ListNode{
				Val: list1.Val,
				Next: mergeTwoLists(list1.Next, list2),
			}
		} else {
			return &ListNode{
				Val: list2.Val,
				Next: mergeTwoLists(list1, list2.Next),
			}
		}
	} else if list2 != nil && list1 == nil {
		return list2
	}

	return list1
}