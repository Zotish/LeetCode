/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
		return head
	}
    dummy := &ListNode{Next: head}
    prev:=dummy

    for head != nil && head.Next != nil {
        nextNode := head.Next
        fmt.Println("NextNode ",nextNode)
        head.Next = nextNode.Next
        fmt.Println("Head.next ",head.Next)
        nextNode.Next = head
        fmt.Println("NextNode.Next ",nextNode.Next)
        prev.Next = nextNode
        fmt.Println("prev>next ",prev.Next)

        prev = head
        fmt.Println("Previous ",prev)
        head = head.Next
        fmt.Println("Head ",head)
    }

    return dummy.Next
}
