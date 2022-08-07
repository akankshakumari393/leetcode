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
    left := head
    right := head
    
    
    for right.Next != nil {
        if right.Next.Val != right.Val {
            left.Next = right.Next
            left = left.Next
        }
        right = right.Next
    }
    left.Next = nil
    return head
}