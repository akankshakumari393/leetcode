/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
    left := head
    right := left
    for right != nil && right.Next != nil {
        left = left.Next
        right = right.Next.Next
        if left == right {
            left2 := head
    
            for left2 != left {
                left = left.Next
                left2 = left2.Next
            }
            return left
        }
    }
    

    return nil
    
}

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


//                     if (fast == slow){
//                         ListNode slow2 = head; 
//                         while (slow2 != slow){
//                             slow = slow.next;
//                             slow2 = slow2.next;
//                         }
//                         return slow;
//                     }