/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func diameterOfBinaryTree(root *TreeNode) int {
    var mx float64
    mx = 0
    heightOfTree(root, &mx)
    return int(mx)
}

func heightOfTree(root *TreeNode, max *float64) float64 {
    if root == nil {
        return 0
    }
    left := heightOfTree(root.Left, max)
    right := heightOfTree(root.Right, max)
    *max = math.Max(*max, left+right)
    return 1 + math.Max(left, right)  
}