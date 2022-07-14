/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

/**
* Example -
* 
* inorder [4, 8, 2, 5, 1, 6, 3, 7]
* postorder [8, 4, 5, 2, 6, 7, 3, 1]
* 
* initially index is 4
*/
func buildTree(inorder []int, postorder []int) *TreeNode {
       if len(postorder) == 0 {
        return nil
    }
    index := searchElement(postorder[len(postorder)-1], inorder)
    node := &TreeNode{
        Val: postorder[len(postorder)-1],
        Left: buildTree(inorder[:index], postorder[:index]),
        Right: buildTree(inorder[index+1:], postorder[index:len(postorder)-1]),
    }
    return node
}

func searchElement(val int, inorder []int) int {
    for i, elem := range inorder {
        if elem == val {
            return i
        }
    }
    return 0
}
