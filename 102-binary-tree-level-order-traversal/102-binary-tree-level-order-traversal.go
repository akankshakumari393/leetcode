/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
    var ans [][]int
    addNodetolevel(root, 0, &ans)
    return ans
}

func addNodetolevel(root *TreeNode, level int, ans *[][]int) {
    if root == nil {
        return
    }
    if len(*ans)-1 < level {
        var newLevel []int
        *ans = append(*ans, newLevel)
    }
    (*ans)[level] = append((*ans)[level], root.Val)
    addNodetolevel(root.Left, level+1, ans)
    addNodetolevel(root.Right, level+1, ans)
}