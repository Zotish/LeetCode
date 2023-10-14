/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteMiddle(head *ListNode) *ListNode {
    n:=0
    temp := head
    for temp!=nil{
    temp=temp.Next
    n++
    }
    middleNode:=n/2
   if head == nil {
        return nil
    }
   temp=head

    if middleNode == 0 {
        head = temp.Next
        return head
    }

    for i := 0; temp != nil && i < middleNode-1; i++ {
        temp = temp.Next
    }

    if temp == nil || temp.Next == nil {
        return nil
    }

    
    next := temp.Next.Next

    temp.Next = next
    return head
}