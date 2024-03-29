// recursion   ---- you would get the answer at the end
func combinationSum(candidates []int, target int) [][]int {
    var twoD [][]int
    var a []int
//    findCombinations(0, target, a, candidates, &twoD)
    backtrack(0, target, a, candidates, &twoD)
    return twoD
}

// Using Recursion
func findCombinations(pos int, target int, ans []int, candidates []int, finalAns *[][]int) {
    // fmt.Println("making call", pos, target, ans, candidates, finalAns)
    if pos == len(candidates) {
        if target == 0 {
            // create new slice
            ans1 := make([]int, len(ans))
            copy(ans1, ans)
            *finalAns = append(*finalAns, ans1)
  //          fmt.Println("Appending the ans to final ans", ans1, finalAns)
        }
        return
    }
    
    // pick the position
    if candidates[pos] <= target {
        // add the element to answer
        //fmt.Println("Adding element to ans", candidates[pos], pos, target, ans, finalAns)        
        ans = append(ans, candidates[pos])
        //fmt.Println("Added element to ans", candidates[pos], pos, target, ans, finalAns)        
        findCombinations(pos, target-candidates[pos], ans, candidates, finalAns)
        //remove the element once tracing is done this represent that even after adding the ele
        //fmt.Println("Removing element from ans", candidates[pos], pos, target, ans, finalAns)        

        ans = ans[:len(ans)-1]
    }
    // not pick the position
    findCombinations(pos+1, target, ans, candidates, finalAns)
}


//backtrack --- answer at every node
func backtrack(pos int, target int, entry []int, candidates []int, res *[][]int) {
    if target < 0 {
        return
    }
    if target == 0 {
        cpy := make([]int, len(entry))
        copy(cpy, entry)
        *res = append(*res, cpy)        
    }
    for i := pos; i < len(candidates); i++ {
        entry = append(entry, candidates[i])
        // because the same number could be repeated so use i
        backtrack(i, target-candidates[i], entry, candidates, res)
        entry = (entry)[:len(entry)-1]
    }
}