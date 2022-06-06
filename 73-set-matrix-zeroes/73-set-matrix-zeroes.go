func setZeroes(matrix [][]int)  {
    m := len(matrix)
	if m == 0 {
		return
	}
	n := len(matrix[0])
    //using extra constant space
    rowToSet0 := make([]int,m)
    columnToSet0 := make([]int,n)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
                rowToSet0[i] = 1
                columnToSet0[j] = 1
            }
		}
	}
    
    
  //   fmt.Println(rowToSet0)   
  //   fmt.Println(columnToSet0)   
  // //  1,2,3
    
    for j := 0; j < len(columnToSet0); j++ {
        if columnToSet0[j] == 1 {
                    for i:=0; i< m;i++{
            matrix[i][j] = 0
        }

        } 
    }

	// for i := 0; i < m; i++ {
	// 	for j := 0; j < n; j++ {
	// fmt.Print(matrix[i][j])
	// }
	// fmt.Println("")
	// }

    
    for i := 0; i < len(rowToSet0); i++ {
        if rowToSet0[i] ==1 {
        for j:=0; j< n;j++{
            matrix[i][j] = 0
        }
        }
    }
    
}