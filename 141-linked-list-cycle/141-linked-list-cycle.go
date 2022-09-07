/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
    left := head
    right := left
    for right != nil && right.Next != nil {
        left = left.Next
        right = right.Next.Next
        if left == right {
            return true
        }
    }
    return false
}