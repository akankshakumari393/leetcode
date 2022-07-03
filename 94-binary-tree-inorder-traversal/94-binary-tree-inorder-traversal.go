/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
    var ans []int
    inorder(root, &ans)
    return ans
}

func inorder(root *TreeNode, ans *[]int) {
    if root == nil {
        return
    }
    if root.Left != nil {
        inorder(root.Left, ans)
    }
    *ans = append(*ans, root.Val)
    if root.Right != nil {
        inorder(root.Right, ans)
    }
}