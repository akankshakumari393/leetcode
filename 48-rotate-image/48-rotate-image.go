func rotate(matrix [][]int)  {
    // transpose of a matrix
    
    n := len(matrix) 
    for i:=0; i<n ;i=i+1 {
        for j:=0; j<i; j=j+1 {
            matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
        }
    }
    // reverse each row 
    for i:=0; i<n ;i=i+1 {
        for j:=0; j<n/2; j=j+1 {
            matrix[i][j], matrix[i][n-1-j] = matrix[i][n-1-j], matrix[i][j]
        }
    }

    
}