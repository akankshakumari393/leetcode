/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
    if head == nil {
        return nil
    }
    lower := head
    right := head
    
    
    for right.Next != nil {
        if right.Next.Val != right.Val {
            lower.Next = right.Next
            lower = lower.Next
        }
        right = right.Next
    }
    lower.Next = nil
    return head
}