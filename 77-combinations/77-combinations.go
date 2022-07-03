func combine(n int, k int) [][]int {
    var result [][]int
    candidates := []int{}
	for i := 1; i <= n; i++ {
		candidates = append(candidates, i)
	}
    findCombination([]int{}, k, candidates, &result)
    return result    
}

// T: O(n!)
// S: O(n!)
func findCombination(included []int, k int, left []int, result *[][]int) {
    if len(included) == k {
        ans1 := append([]int{}, included...)
        *result = append(*result, ans1)
        return
    }
    for idx, l := range left{
        findCombination(
            append(included, l),
            k,
            append([]int{}, left[idx+1:]...),
            result)        
    }
    
}
