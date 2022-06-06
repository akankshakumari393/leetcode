func setZeroes(matrix [][]int)  {
    m := len(matrix)
	if m == 0 {
		return
	}
	n := len(matrix[0])
    //using extra constant space
    clearRow := make([]int,m)
    clearColumn := make([]int,n)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
                clearRow[i] = 1
                clearColumn[j] = 1
            }
		}
	}
    
    
  //   fmt.Println(rowToSet0)   
  //   fmt.Println(columnToSet0)   
  // //  1,2,3
    
//     for j := 0; j < len(clearcolumn); j++ {
//         if clearColumn[j] == 1 {
//                     for i:=0; i< m;i++{
//             matrix[i][j] = 0
//         }

//         } 
//     }

// 	// for i := 0; i < m; i++ {
// 	// 	for j := 0; j < n; j++ {
// 	// fmt.Print(matrix[i][j])
// 	// }
// 	// fmt.Println("")
// 	// }

    
//     for i := 0; i < len(clearRow); i++ {
//         if clearRow[i] ==1 {
//         for j:=0; j< n;j++{
//             matrix[i][j] = 0
//         }
//         }
//     }
    
    // optimized 

	for col, val := range clearColumn {
		if val == 1 {
			for i := 0; i < m; i++ {
				matrix[i][col] = 0
			}
		}
	}
	for row, val := range clearRow {
		if val == 1 {
			for i := 0; i < n; i++ {
				matrix[row][i] = 0
			}
		}
	}
    
    
    
    
}