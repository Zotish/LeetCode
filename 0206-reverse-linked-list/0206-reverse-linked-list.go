/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
    var prev *ListNode
    current := head
    for current != nil {
        nextTemp:= current.Next // store next node
        current.Next = prev      // reverse current node's direction
        prev = current           // move prev to current node
        current = nextTemp       // move to the next node
    }
    return prev
}