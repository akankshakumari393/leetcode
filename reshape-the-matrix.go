func matrixReshape(mat [][]int, r int, c int) [][]int {

	// check if number of elements in input is equal to given r*c
	m := len(mat)
	n := len(mat[0])
    if (r*c != m*n) || (r == m && c == n) {
		// if number of elements are not equal or if rows and columns are equal return original matrix
		return mat
	}
	result := make([][]int, 0)
	for i := 0; i < r; i++ {
		result = append(result, make([]int, c))
	}
	for i := 0; i < r*c; i++ {
		result[i/c][i%c] = mat[i/n][i%n]
	}
	return result
}
