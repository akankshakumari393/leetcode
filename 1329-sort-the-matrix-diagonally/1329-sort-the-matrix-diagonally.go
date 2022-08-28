func diagonalSort(mat [][]int) [][]int {
    m:=len(mat)
    n:=len(mat[0])
    // col 0
    for i:=0; i<m; i++{
        CountSort(mat, i, 0, m, n)
    }
    // row 0
    for j:=1 ;j<n; j++{
        CountSort(mat, 0, j, m, n)
    }
    return mat
}


func Sort(mat [][]int, row, col, m, n int) {
    temp:=make([]int, 0)
    x:= row
    y:= col
    for x<m&&y<n {
        temp = append(temp, mat[x][y])
        x++
        y++
    }
    sort.Ints(temp)
    idx:=0
    x=row
    y=col
    for x<m&&y<n {
        mat[x][y] = temp[idx]
        idx++
        x++
        y++
    }    
}

func CountSort(mat [][]int, row, col, m, n int) {
    temp:=make([]int, 101)
    x:= row
    y:= col
    for x<m&&y<n{
        temp[mat[x][y]]++
        x++
        y++
    }
    x=row
    y=col

    for i:=1; i<101; i++{
        for temp[i] > 0 {
            mat[x][y] = i
            temp[i]--
            x++
            y++
        }
    }
}
