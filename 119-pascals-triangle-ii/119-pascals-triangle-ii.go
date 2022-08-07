// val = val * (rowIndex - j)/(j + 1)
// At rowIndex there are rowIndex + 1 elements
func getRow(rowIndex int) []int {
    var res []int
    val := 1
    for j := 0; j <= rowIndex; j++ {
        res = append(res, val)
        val = val * (rowIndex - j) / (j + 1)
    }
    return res
}


// for 4th index

// 1 = 1
// 1 * (4-0)/(0+1)= 1*4/1 = 4
// 4 * (4-1)/(1+1)= 4*3/2 = 6
// 6 * (4-2)/(2+1)= 6*2/3 = 4
// 4 * (4-3)/(3+1)= 4*1/4 = 1
