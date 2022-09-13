func combinationSum2(candidates []int, target int) [][]int {
    var twoD [][]int
    var a []int
    sort.Ints(candidates)
    //findCombinations(0, target, a, candidates, &twoD)
    backtrack(0, target, a, candidates, &twoD)
    return twoD
}

// to solve this using recursion we have to maintain that the solution 2d doesn't contain same list
func findCombinations(pos int, target int, entry []int, candidates []int, res *[][]int) {
    fmt.Println("recursion called for with ", entry)
    if pos == len(candidates) {
        if target == 0 {
            // create new slice
            cpy := make([]int, len(entry))
            copy(cpy, entry)
            *res = append(*res, cpy)
  //          fmt.Println("Appending the ans to final ans", ans1, finalAns)
        }
        return        
    }
    // pick the position
    if candidates[pos] <= target {
        // add the element to answer
        //fmt.Println("Adding element to ans", candidates[pos], pos, target, ans, finalAns)        
        entry = append(entry, candidates[pos])
        // pos +1 as we do not want an element to repeat itself        
        findCombinations(pos+1, target-candidates[pos], entry, candidates, res)
        //remove the element once tracing is done this represent that even after adding the ele
        //fmt.Println("Removing element from ans", candidates[pos], pos, target, ans, finalAns)        

        entry = entry[:len(entry)-1]
    }
    // not pick the position
    findCombinations(pos+1, target, entry, candidates, res)
}



//backtrack --- answer at every node
func backtrack(pos int, target int, entry []int, candidates []int, res *[][]int) {
    //fmt.Println("backtrack called for with ", entry)
    if target < 0 {
        return
    }
    if target == 0 {
        cpy := make([]int, len(entry))
        copy(cpy, entry)
        *res = append(*res, cpy)        
    }
    for i := pos; i < len(candidates); i++ {
        // backtracking to avoid duplication
       if i > pos && candidates[i] == candidates[i - 1]{
           // fmt.Println("returning as i is ", i ,"pos ", pos, "entry ", entry)
            continue
        }            

        
        
        
        entry = append(entry, candidates[i])
        // because a number should not repeat itself so use i+1
        backtrack(i+1, target-candidates[i], entry, candidates, res)
        entry = (entry)[:len(entry)-1]
    }
}