/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	index := search(preorder[0], inorder)
	node := &TreeNode{
		Val:   preorder[0],
		Left:  buildTree(preorder[1:index+1], inorder[:index]),
		Right: buildTree(preorder[index+1:], inorder[index+1:]),
	}

	return node
}

func search(val int, inorder []int) int {
	for i := 0; i < len(inorder); i++ {
		if val == inorder[i] {
			return i
		}
	}
	return 0
}