/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedArrayToBST(nums []int) *TreeNode {
    // find the mid element and make it root
    mid := (len(nums)-1)/2
    if mid < 0 || mid >= len(nums){
        return nil
    }
    node := &TreeNode{}
    node.Val = nums[mid]
    // Do the same for Right and left array
    node.Right = sortedArrayToBST(nums[mid+1:])
    node.Left = sortedArrayToBST(nums[:mid])
    return node
}