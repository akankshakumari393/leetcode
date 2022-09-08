/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
    curA := headA
	curB := headB
    // go till one of the poiter is nil, this will have the bigger list having extra node
	for curA != nil && curB != nil {
		curA = curA.Next
		curB = curB.Next
	}
    // if curA is bigger then traverse that much so that curA is nil
	for curA != nil {
		headA = headA.Next
		curA = curA.Next
	}
    // if curB is bigger then traverse that much so that curB is nil
	for curB != nil {
		headB = headB.Next
		curB = curB.Next
	}
    // now headA and headB have equal number of element
	for headA != headB {
		headA = headA.Next
		headB = headB.Next
	}

	return headA
}