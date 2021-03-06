func findPaths(m int, n int, maxMove int, startRow int, startColumn int) int {
    //return BruteForceFindPaths(m, n, maxMove, startRow, startColumn)
    return Dynamic(m, n, maxMove, startRow, startColumn)
}

// Time complexity 4^n - size of recursion tree
// Space complexity - depth of recursion tree - o(n)
func BruteForceFindPaths(m int, n int, N int, i int, j int) int {
    if (i == m || j == n || i < 0 || j < 0) {
        return 1
    }
    if (N == 0) {
     return 0   
    }
    return BruteForceFindPaths(m, n, N - 1, i - 1, j) + BruteForceFindPaths(m, n, N - 1, i + 1, j) + BruteForceFindPaths(m, n, N - 1, i, j - 1) + BruteForceFindPaths(m, n, N - 1, i, j + 1)
}    

func Dynamic(m int, n int, N int, startRow int, startColumn int) int{
    mod := int(math.Pow10(9) + 7)
    dp := make([][][]int, N + 1)
    for grid := range dp {
        dp[grid] = make([][]int, m)
        for row := range dp[grid] {
            dp[grid][row] = make([]int, n)
        }
    }
    dirs := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
    for s := 1; s <= N; s++ {
        for i := 0; i < m; i++ {
            for j := 0; j < n; j++ {
                for _, dir := range dirs {
                    tx := dir[0] + i
                    ty := dir[1] + j
                    if ty < 0 || tx < 0 || ty >= n || tx >= m {
                        dp[s][i][j] += 1
                    } else {
                        dp[s][i][j] = (dp[s][i][j] + dp[s-1][tx][ty]) % mod
                    }
                }
            }
        }
    }
    return dp[N][startRow][startColumn]
}