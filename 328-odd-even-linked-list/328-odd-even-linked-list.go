/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func oddEvenList(head *ListNode) *ListNode {
    if head == nil {
        return nil
    }
    odd := head
    even := head.Next
    evenHead := even
    for odd.Next != nil && even.Next != nil {
        odd.Next = even.Next
        odd = odd.Next
        even.Next = odd.Next
        even = even.Next
    }
    // now we have the last odd node pointing to last even
    // we need to make it point to first even
    odd.Next = evenHead
    return head
}