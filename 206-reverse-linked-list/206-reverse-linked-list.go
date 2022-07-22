/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
    if head == nil{
		return nil
	}
	temp := head.Next
	head.Next = nil
	return reversePair(head, temp)
}

func reversePair(last *ListNode, head *ListNode) *ListNode {
	if head == nil{
		return last
	}
	temp := head.Next
	head.Next = last
	return reversePair(head, temp)
}
