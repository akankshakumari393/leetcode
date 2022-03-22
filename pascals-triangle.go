// Property of pascal triangle

// 1
// 1,1
// 1.2.1

// first and last element of a row is 1 always
// the number of element in each row is rowNumber+1


func generate(numRows int) [][]int {
	result := make([][]int, 0)
	for i := 0; i < numRows; i++ {
		rowElement := make([]int, i+1)
		rowElement[0] = 1 // set first column of every row to 1
	    result = append(result, rowElement) // allot space for each row
		for j := 1; j < i+1; j++ {

			if i == j { // last element is always 1
				result[i][j] = 1
			} else {
				result[i][j] = result[i-1][j] + result[i-1][j-1]
			}
		}
	}
	return result
}
