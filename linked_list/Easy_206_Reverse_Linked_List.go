func reverseList(head *ListNode) *ListNode {
    var prev *ListNode

    for head != nil {
        head.Next, prev, head = prev, head, head.Next
    }
    return prev
}

func reverseList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }

    reversedHead := reverseList(head.Next)
    head.Next.Next = head
    head.Next = nil
    return reversedHead
}
