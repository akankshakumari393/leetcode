func permute(nums []int) [][]int {
    var result [][]int
    findPermutation([]int{},nums,&result)
    return result
}

// // T: O(n!)
// // S: O(n!)
func findPermutation(included []int, left []int, result *[][]int) {
    if len(left) == 0 {
        ans1 := append([]int{}, included...)
        *result = append(*result, ans1)
        return
    }
    for idx, l := range left {
        findPermutation(
            append(included, l),
            append(append([]int{}, left[:idx]...),left[idx+1:]...),
            result)        
    }
    
}
