package main

/*Definition for singly-linked list.*/
//  type ListNode struct {
//      Val int
//      Next *ListNode
//  }

 func deleteDuplicates(head *ListNode) *ListNode {
    if head != nil && head.Next != nil {
        if head.Val == head.Next.Val {
            return deleteDuplicates(head.Next)
        }
		return &ListNode{
			Val: head.Val,
			Next: deleteDuplicates(head.Next),
		}
    }
    return head
}