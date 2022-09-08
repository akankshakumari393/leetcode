/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
    if head == nil {
        return nil
	}
	cloneMap := map[*Node]*Node{}
	for cur := head; cur != nil; cur = cur.Next {
		cloneMap[cur] = &Node{}
	}
	for k, v := range cloneMap {
		v.Val = k.Val
		if k.Next != nil {
			v.Next = cloneMap[k.Next]
		}
		if k.Random != nil {
			v.Random = cloneMap[k.Random]
		}
	}
	return cloneMap[head]
}