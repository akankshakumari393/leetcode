/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    // set this to zero
    var dummy = &ListNode{Val: 0}
	var currentAddress = dummy
	var sum int
	for l1 != nil || l2 != nil {
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
        // add a new node to the list 
		currentAddress.Next = &ListNode{Val: sum % 10}
		currentAddress = currentAddress.Next
		if sum > 9 {
			sum = 1 // carry forward
		} else {
			sum = 0
		}
	}
	if sum == 1 {
		currentAddress.Next = &ListNode{Val: sum}
	}
	return dummy.Next
}