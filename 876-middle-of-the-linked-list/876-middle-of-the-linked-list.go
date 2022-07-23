/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func middleNode(head *ListNode) *ListNode {
    if head == nil {
        return nil
    }
    length := lengthOfLinkedList(head)
    for i:=1; i<=length/2 ;i++ {
        head = head.Next       
    }
    return head
}

func lengthOfLinkedList(head *ListNode) int {
    if head == nil {
        return 0
    }
    return 1 + lengthOfLinkedList(head.Next)
}