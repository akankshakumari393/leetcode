func gameOfLife(board [][]int)  {
    dx:= []int{1, 1, 0, -1,-1, -1, 0, 1}
    dy:= []int{0, 1, 1,  1, 0, -1, -1,-1}
    R := len(board)
    C := len(board[0])
    for i:=0 ; i< R; i++ {
        for j:=0; j< C; j++ {
            liveCount := 0
            for k:=0; k<8; k++ {
                if IsSafe(i+dx[k], j+ dy[k], R, C) && int(math.Abs(float64(board[i+dx[k]][j+dy[k]]))) == 1 {
                    liveCount++                    
                }            
            }
            if board[i][j] == 0 && liveCount == 3 {
                board[i][j]= 2
            } else if board[i][j] == 1 && liveCount < 2  {
                board[i][j]= -1
            } else if board[i][j] == 1 && liveCount > 3 {
                board[i][j]= -1
            }
        }
    }

    // for i:=0 ; i< len(board); i++ {
    //     for j:=0; j< len(board[0]); j++{
    //         fmt.Print(" ", board[i][j])
    //     }
    //     fmt.Println("")
    // }
   
    for i:=0 ; i< R; i++ {
        for j:=0; j< C; j++{
            if board[i][j] < 0 {
                board[i][j] = 0
            }else if board[i][j] > 1 {
                board[i][j] = 1
            }
        }
    }
}

func IsSafe(x int, y int, R int, C int) bool{
    if x>=0 && x < R && y<C && y >=0 {
        return true
    }
    return false
}