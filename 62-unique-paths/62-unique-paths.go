func uniquePaths(m int, n int) int {
    return uniquePathsDP(m, n)
    //return countPaths(0,0,m,n) 
}
// recursive approach
func countPaths(i int,j int,m int,n int) int {
    if(i==(m-1)&&j==(n-1)) {
     return 1   
    }
    if(i>=m||j>=n) {
      return 0  
    }
   return countPaths(i+1,j,m,n)+countPaths(i,j+1,m,n)
}


// dynamic approach
// DP[m-1][n-1] = 1
// DP[m-1][0...n-2] = 1
// DP[0...m-2][n-1] = 1

func uniquePathsDP(m int, n int) int {
    dp := make([][]int, m) // initialize a slice of m elements
    for idx, _ := range dp {
        dp[idx] = make([]int, n) // initialize each row
        dp[idx][n-1] = 1 //initailize the last column as 1
    }
    for idx, _ := range dp[0] {
        dp[m-1][idx] = 1  // intialize last row as 1
    }
    for i:=m-2; i >= 0; i-- {
        for j:=n-2; j >= 0; j-- {
            dp[i][j] = dp[i+1][j] + dp[i][j+1]
        }
    }
    return dp[0][0]
}