/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
    if head == nil {
        return nil
    }
    start := &ListNode{Next : head}
    slow := start
    fast := start
    for i := 1; i <= n; i++ {
        fast = fast.Next
    }
    // fast id pointing to nth node
    
    for fast.Next != nil {
        fast = fast.Next
        slow = slow.Next
    }
    // slow is at one position previous from n-1 th node from beginning we need to delete node next to slow
    slow.Next = slow.Next.Next
    // head is always going to point to first node
    return start.Next
}