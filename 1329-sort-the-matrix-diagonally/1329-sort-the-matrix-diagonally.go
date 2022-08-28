func diagonalSort(mat [][]int) [][]int {
    m:=len(mat)
    n:=len(mat[0])
    // col 0
    for i:=0; i<m; i++{
        Sort(mat, i, 0, m, n)
    }
    // row 0
    for j:=1 ;j<n; j++{
        Sort(mat, 0, j, m, n)
    }
    return mat
}


func Sort(mat [][]int, row, col, m, n int) {
        temp:=make([]int, 0)
        x:= row
        y:= col
        for x<m&&y<n{
            temp = append(temp, mat[x][y])
            x++
            y++
        }
        sort.Ints(temp)
        idx:=0
        x=row
        y=col
        for x<m&&y<n{
            mat[x][y] = temp[idx]
            idx++
            x++
            y++
        }    
}