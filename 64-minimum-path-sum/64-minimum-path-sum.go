func minPathSum(grid [][]int) int {
    var ans [][]int
    m := len(grid)
    n := len(grid[0])
    for i:=0; i< m; i++ {
    ans = append(ans, make([]int, n))
    }
    ans[0][0] = grid[0][0]
    
    for i:=1; i< m; i++ {
        ans[i][0] = grid[i][0] + ans[i-1][0]
    }
    
    for j:=1; j< n; j++ {
        ans[0][j] = grid[0][j] + ans[0][j-1]
    }
    
    
    for i:= 1; i< m; i++{
        for j:= 1; j< n; j++ {
            ans[i][j] = grid[i][j] + int((math.Min(float64(ans[i-1][j]), float64(ans[i][j-1]))))                
        }
    }
    return ans[m-1][n-1]                            
}