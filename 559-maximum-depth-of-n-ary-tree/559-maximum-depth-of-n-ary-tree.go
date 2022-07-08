/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func maxDepth(root *Node) int {
    if root == nil {
        return 0
    }
    depth := 0
    for _, node := range root.Children {
        d := maxDepth(node)
        if d > depth {
            depth = d
        }
    }
    return depth + 1
}