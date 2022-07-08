/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

func cloneGraph(node *Node) *Node {
    if node == nil {
        return nil
    }
    return clone(node, make(map[int]*Node))    
}
func clone(node *Node, hm map[int]*Node) *Node {
    if node, ok := hm[node.Val]; ok {
            return node        
    } 
    newNode := &Node{}
    newNode.Val = node.Val
    hm[node.Val] = newNode
    for _, neighbor := range node.Neighbors {
        newNode.Neighbors = append(newNode.Neighbors, clone(neighbor, hm))   
    }
    return newNode
}