/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rightSideView(root *TreeNode) []int {
	result := []int{}
	levelOrder(root, &result, 0)
	return result
}

func levelOrder(current *TreeNode, result *[]int, level int) {
	if current == nil {
		return
	}

	if level == len(*result) {
		*result = append(*result, current.Val)
	}

	(*result)[level] = current.Val
	levelOrder(current.Left, result, level+1)
	levelOrder(current.Right, result, level+1)
}