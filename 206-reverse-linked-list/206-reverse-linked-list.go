/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
    if head == nil {
        return nil
    }
    var prev *ListNode // this needs to be nil
    curr := head
    for curr != nil {
        // Store next
       next := curr.Next
       curr.Next = prev
       prev = curr
       curr = next
    }
    head = prev
    return head
}