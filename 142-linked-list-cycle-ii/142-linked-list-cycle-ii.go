/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
    slow := head
    fast := slow
    for fast != nil && fast.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
        if fast == slow {
            slow2 := head
    
            for slow2 != slow {
                slow = slow.Next
                slow2 = slow2.Next
            }
            return slow
        }
    }
    return nil
    
}