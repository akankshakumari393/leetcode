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

func Dynamic(m int, n int, N int, i int, j int) int{
    mod := int(math.Pow10(9) + 7)
    dp := make([][][]int, N + 1)
    for i := range dp {
        dp[i] = make([][]int, m)
        for I := range dp[i] {
            dp[i][I] = make([]int, n)
        }
    }
    dirs := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
    for s := 1; s <= N; s++ {
        for x := 0; x < m; x++ {
            for y := 0; y < n; y++ {
                for _, dir := range dirs {
                    tx := dir[0] + x
                    ty := dir[1] + y
                    if ty < 0 || tx < 0 || ty >= n || tx >= m {
                        dp[s][x][y] += 1
                    } else {
                        dp[s][x][y] = (dp[s][x][y] + dp[s-1][tx][ty]) % mod
                    }
                }
            }
        }
    }
    return dp[N][i][j]
}