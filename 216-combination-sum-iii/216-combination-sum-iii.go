func combinationSum3(k int, n int) [][]int {
    var twoD [][]int
    var a []int
    candidates := make([]int, 9)
    for i := 1; i<=9; i++ {
        candidates[i-1] = i
    }
    
    //findCombinations(0, target, a, candidates, &twoD)
    backtrack(0, n, a, candidates, &twoD, k)
    return twoD
}

//backtrack --- answer at every node
func backtrack(pos int, target int, entry []int, candidates []int, res *[][]int, k int) {
 //   fmt.Println("backtrack called for with ", entry)
    if target < 0 {
        return
    }
    if target == 0 && len(entry) == k {
        cpy := make([]int, len(entry))
        copy(cpy, entry)
        *res = append(*res, cpy)        
    }
    for i := pos; i < len(candidates); i++ {
        // backtracking to avoid duplication. This is not required here as the candidates don't have duplicate values
       // if i > pos && candidates[i] == candidates[i - 1]{
       //     // fmt.Println("returning as i is ", i ,"pos ", pos, "entry ", entry)
       //      continue
       //  }
        
        entry = append(entry, candidates[i])
        // because a number should not repeat itself so use i+1
        backtrack(i+1, target-candidates[i], entry, candidates, res, k)
        entry = (entry)[:len(entry)-1]
    }
}