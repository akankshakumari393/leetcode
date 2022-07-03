/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {
    return isValid(root, math.MinInt, math.MaxInt)
}

func isValid(root *TreeNode, min int, max int) bool {
    if root == nil {
        return true
    }
    if root.Val <= min {
        return false
    }
    if root.Val >= max {
        return false
    }
    
    if isValid(root.Left, min, root.Val) && isValid(root.Right, root.Val, max) {
        return true
    }
    return false
}