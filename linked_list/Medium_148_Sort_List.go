func sortList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    slow, fast := head, head.Next
    for fast != nil && fast.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
    }
    right := sortList(slow.Next)
    slow.Next = nil
    left := sortList(head) 

    // Merge lists
    sentinel := new(ListNode)
    cur := sentinel

    for left != nil && right != nil {
        if left.Val < right.Val {
            cur.Next = left
            left = left.Next
        } else {
            cur.Next = right
            right = right.Next
        }
        cur = cur.Next
    }
    if left != nil {
        cur.Next = left
    } else {
        cur.Next = right
    }
    return sentinel.Next
}
