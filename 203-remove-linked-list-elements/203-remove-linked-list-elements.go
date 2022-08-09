/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements(head *ListNode, val int) *ListNode {
	left := &ListNode{Next: head}
	right := left
	for right.Next != nil {
		if right.Next.Val == val {
			right.Next = right.Next.Next
		} else {
			right = right.Next
		}
	}

	return left.Next
}