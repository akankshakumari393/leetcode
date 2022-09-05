/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func middleNode(head *ListNode) *ListNode {
    left := head
    right := left
    for right != nil && right.Next != nil {
        left = left.Next
        right = right.Next.Next
    }
    return left
}