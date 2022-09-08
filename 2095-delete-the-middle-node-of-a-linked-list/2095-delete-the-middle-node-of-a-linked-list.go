/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteMiddle(head *ListNode) *ListNode {
    if head == nil {
        return nil
    }
    dummy := &ListNode{Next : head}
    // pointing slow to dummy beacuse this would allow us to delete the node
    slow := dummy
    fast := dummy
    for fast.Next != nil && fast.Next.Next != nil {
        fast = fast.Next.Next
        slow = slow.Next
    }
    slow.Next = slow.Next.Next
    return dummy.Next
}