func removeElements(head *ListNode, val int) *ListNode {
    var res *ListNode

    // find head 
    for head != nil {
        if head.Val == val {
            head = head.Next
        } else {
            res = head
            break
        }
    }

    current := res
    for current != nil && current.Next != nil {
        if current.Next.Val == val {
            current.Next = current.Next.Next
        } else {
            current = current.Next
        }
    }
    return res
}