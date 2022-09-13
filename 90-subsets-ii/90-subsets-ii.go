func subsetsWithDup(nums []int) [][]int {
    var res [][]int
    var entry []int
    sort.Ints(nums)
    backtrack(0, nums, entry, &res)
    return res

}

func backtrack(pos int, candidates []int, entry []int, res *[][]int) {
    cpy := make([]int, len(entry))
    copy(cpy, entry)
    *res = append(*res, cpy)
    
    for i := pos; i < len(candidates); i++ {
        // the solution set must not contain same list, that's why continue in case of same subtree 
        if i > pos && candidates[i] == candidates[i-1] {
            continue
        }
        
        entry = append(entry, candidates[i])
        // the number should not repeat itself that's why i+1
        backtrack(i+1, candidates, entry, res)
        entry = (entry)[:len(entry)-1]
    }
}