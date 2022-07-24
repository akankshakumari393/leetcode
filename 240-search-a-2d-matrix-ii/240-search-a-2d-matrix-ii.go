// Time Complexity O(m+n)
// Space Complexity O(1)
func searchMatrix(matrix [][]int, target int) bool {
    return RecursiveSolution(matrix, 0, len(matrix[0])-1, target) 
}

// Start from top right element
func RecursiveSolution(matrix [][]int, m, n, target int) bool {
    if m >= len(matrix) || m < 0 || n >= len(matrix[0]) || n < 0 {
        return false
    }
    if target == matrix[m][n] {
        return true
    }
    // if target > the current element search left
    if target >= matrix[m][n] {
     return RecursiveSolution(matrix, m+1, n, target)   
    }
    // if target < the current element search down
    return RecursiveSolution(matrix, m, n-1, target)
}
