func solve(board [][]byte)  {
    if len(board) == 0 || len(board[0]) == 0 {
      return  
    } 
    if len(board) < 3 || len(board[0]) < 3 {
        return
    }
    m := len(board)
    n := len(board[0])
    // traverse over first column and last column
    for i := 0; i < m; i++ {
        if (board[i][0] == 'O') {
         dfs(board, i, 0)   
        }
        if (board[i][n - 1] == 'O'){
          dfs(board, i, n - 1)  
        }
    }
    // traverse over first row and last row
    for j := 1; j < n - 1; j++ {
        if (board[0][j] == 'O'){
         dfs(board, 0, j)   
        }
        if (board[m - 1][j] == 'O') {
         dfs(board, m - 1, j)   
        }
    }
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if (board[i][j] == 'O') {
             board[i][j] = 'X'   
            }
            if (board[i][j] == '*') {
              board[i][j] = 'O'        
            }
        }
    }

    
}


func dfs(board [][]byte , r int, c int) {
    if (r < 0 || c < 0 || r > len(board) - 1 || c > len(board[0]) - 1 || board[r][c] != 'O') {
        return
    }
    board[r][c] = '*';
    dfs(board, r + 1, c);
    dfs(board, r - 1, c);
    dfs(board, r, c + 1);
    dfs(board, r, c - 1);
}