/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
    fast := head
    slow := head
    for fast.Next != nil && fast.Next.Next != nil {
        fast = fast.Next.Next
        slow = slow.Next
    }
    // slow is the middle node
    // reverse iteratively from next of middle node
    
    var prev *ListNode
    curr := slow.Next
    next := &ListNode{}
    for curr != nil {
       next = curr.Next
       curr.Next = prev
       prev = curr
       curr = next
    }
    lasthead := prev
    // situation after reverse would be two seperate linkedlist
    // reset
    fast = head
    for lasthead != nil {
        if fast.Val != lasthead.Val {
            return false
        }
        lasthead = lasthead.Next
        fast = fast.Next
    }
    return true
    
}