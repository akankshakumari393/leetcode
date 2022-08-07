/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func postorderTraversal(root *TreeNode) []int {
    if root == nil {
        return nil
    }
    var ans []int
    traverse(root, &ans)
    return ans
}

func traverse(root *TreeNode, ans *[]int) {
    if root == nil {
        return
    }
    
    traverse(root.Left, ans)
    traverse(root.Right, ans)
    *ans = append(*ans, root.Val)
}