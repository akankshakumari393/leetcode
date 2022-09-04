/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func recurse(root *TreeNode, ans *int, maxValueSeen int) {
	if root == nil {
		return
	}
	if root.Val >= maxValueSeen {
		*ans++
        maxValueSeen = root.Val
	}
	recurse(root.Left, ans, maxValueSeen)
	recurse(root.Right, ans, maxValueSeen)
}

func goodNodes(root *TreeNode) int {
	ans := 0
	recurse(root, &ans, root.Val)
	return ans
}