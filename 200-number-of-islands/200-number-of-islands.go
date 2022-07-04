func numIslands(grid [][]byte) int {
    var island int
    visitedElement := make(map[string]bool)
    
    for row := 0; row < len(grid); row++ {
        for column := 0; column < len(grid[row]); column++ {
            island += checkIsland(visitedElement, grid, row, column)
        }
    }
    return island
}

func checkIsland(visitedElement map[string]bool, grid [][]byte, row, column int) int {
    if _, isVisited := visitedElement[key(row, column)]; isVisited {
        return 0
    }
    
    if row < 0 || row >= len(grid) || column < 0 || column >= len(grid[0]) {
        return 0
    }
    
    visitedElement[key(row, column)] = true
   
    if grid[row][column] == '0' {
        return 0
    }
    
    checkIsland(visitedElement, grid, row, column+1)
    checkIsland(visitedElement, grid, row, column-1)
    checkIsland(visitedElement, grid, row+1, column)
    checkIsland(visitedElement, grid, row-1, column)
    
    return 1
}

func key(row, column int) string {
    return fmt.Sprintf("%d,%d", row, column)
}