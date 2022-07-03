/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
    if head == nil ||head.Next == nil || k== 0 {
        return head
    }
    // calculate length
    cur := head
    count := 1
    for cur.Next != nil {
        count++
        cur = cur.Next
    }

    if k >= count {
        k= k % count
    }
    fmt.Println(k)
    k = count - k
    
    // creates a circular linked list
    cur.Next = head

     for ;k>=1;k-- {
        cur = cur.Next
    }
    head = cur.Next
    cur.Next = nil  
    
    return head
    
}