/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, left int, right int) *ListNode {
    if head == nil || left==right {
        return head
    }
    
//   1 -> 2 -> 3 -> 4 -> 5 -> 6 
    dummy := &ListNode{
        Val : -1,
    }
    // to handle cases if left and riht are first and last node
    dummy.Next = head
    prev := dummy
    for i:=1; i <= left-1; i++ {
        prev = prev.Next
    }
//  1 -> 2 -> 3 -> 4 -> 5 -> 6
// head,prev pointing at 1
    
    tail := prev.Next
//  1 -> 2 -> 3 -> 4 -> 5 -> 6
// head,prev pointing at 1
// tail pointing at 2
    
    for i:= 1; i <= right-left; i++ {
        tmp := prev.Next
        prev.Next = tail.Next
        tail.Next = tail.Next.Next
        prev.Next.Next = tmp
    }
    
    return dummy.Next
}