/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }

    length := lengthOfLinkedList(head)

    dummyHead := &ListNode{}
    dummyHead.Next = head

    pre := dummyHead
    var cur, nex *ListNode

    for length >= k {
        cur = pre.Next
        nex = cur.Next
        for i := 1; i < k; i++ {
            cur.Next = nex.Next
            nex.Next = pre.Next
            pre.Next = nex
            nex = cur.Next
        }
        pre = cur
        length -= k
    }
    return dummyHead.Next
}

func lengthOfLinkedList(head *ListNode) int {
    length := 0
    for head != nil {
        length++
        head = head.Next
    }
    return length
}
