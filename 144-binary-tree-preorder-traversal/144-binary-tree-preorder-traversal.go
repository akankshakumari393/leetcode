func preorderTraversal(root *TreeNode) []int {
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
    
    *ans = append(*ans, root.Val)
    traverse(root.Left, ans)
    traverse(root.Right, ans)
}