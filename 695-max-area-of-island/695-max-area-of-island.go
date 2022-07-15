func maxAreaOfIsland(grid [][]int) int {
    R := len(grid)
    C := len(grid[0])
    max := 0
    for i:=0; i< R; i++ {
        for j:=0; j<C; j++ {
            if grid[i][j] == 1 {
                area := traverse(i, j, R, C, grid)
                max = maximum(max, area)                
            }
        }
    }
    return max
}

func traverse(i, j, R, C int, grid [][]int) int {
    if i >= R || j >=C || i < 0 || j < 0 {
        return 0
    }
    if grid[i][j] == 0 || grid[i][j] == 2{
        return 0
    }
    // mark the grid visited
    grid[i][j] = 2
    return 1+traverse(i+1, j, R, C, grid)+traverse(i, j+1, R, C, grid)+traverse(i-1, j, R, C, grid)+traverse(i, j-1, R, C, grid)
}

func maximum(x, y int) int {
    if x > y {
        return x
    } 
    return y
}