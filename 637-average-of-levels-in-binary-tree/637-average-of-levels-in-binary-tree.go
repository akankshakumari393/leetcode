/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func averageOfLevels(root *TreeNode) []float64 {
    result := []float64{}
    if root  == nil{
        return []float64{}
    }
    queue := []*TreeNode{root}
    for len(queue) > 0{
        temp := []*TreeNode{}
        sum := 0
        for _, node := range queue{
            sum += node.Val
            if node.Left != nil{
                temp = append(temp, node.Left)
            }
            if node.Right != nil{
                temp = append(temp, node.Right)
            }
            
        }
        result = append(result,float64(sum)/float64(len(queue)))
        queue = temp
    }
    return result
}