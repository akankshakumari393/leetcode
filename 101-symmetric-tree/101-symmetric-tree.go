/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
    if root == nil || IsSymmetricTree(root.Left,root.Right) {
        return true
    }
    return false
}
    // For two trees to be mirror images, the following
    // three conditions must be true
    // 1.) Their root node's key must be same
    // 2.) left subtree of left tree and right subtree of
    // right tree have to be mirror images
    // 3.) right subtree of left tree and left subtree of
    // right tree have to be mirror images
func IsSymmetricTree(p *TreeNode, q *TreeNode) bool {
    if p== nil && q == nil{
        return true
    }
    if p!= nil && q!= nil && p.Val== q.Val && IsSymmetricTree(p.Left, q.Right) && IsSymmetricTree(p.Right,q.Left) {
        return true
    }
    return false
}