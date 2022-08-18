/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
 	dummy := &ListNode{}
	prev := dummy
	for list1 != nil || list2 != nil {
		if list1 == nil {
			prev.Next = list2
			break
		}

		if list2 == nil {
			prev.Next = list1
			break
		}

		if list1.Val < list2.Val {
			prev.Next = list1
			list1 = list1.Next
		} else {
			prev.Next = list2
			list2 = list2.Next
		}

		prev = prev.Next
	}

	return dummy.Next   
}