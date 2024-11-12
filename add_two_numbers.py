# Definition for singly-linked list.
from typing import Optional


class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next

class Solution:
    @staticmethod
    def add_two_numbers(l1: Optional[ListNode], l2: Optional[ListNode]) -> Optional[ListNode]:
        total, carry = 0, 0
        dummy = ListNode()
        result = dummy
        while l1 or l2 or carry:
            total = carry
            if l1:
                total += l1.val
                l1 = l1.next

            if l2:
                total += l2.val
                l2 = l2.next

            final = total % 10
            carry = total // 10

            dummy.next = ListNode(final)
            dummy = dummy.next

        return result.next



if __name__ == '__main__':
    l_one = ListNode(val=2, next=ListNode(4, ListNode(val=3)))
    l_two = ListNode(val=5, next=ListNode(6, ListNode(val=4)))
    print(Solution.add_two_numbers(l_one, l_two).val)
