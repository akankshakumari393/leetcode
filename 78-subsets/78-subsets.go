// func subsets(nums []int) [][]int {
//     // number of paranthesis left to be printed
// 	res := [][]int{}
// 	powerset(&res, []int{}, nums, 0)
// 	return res
// }

func powerset(res *[][]int, currentSet []int, nums []int, level int){
    if level == len(nums) {
        *res = append(*res, currentSet)
        currentSet = []int{}
        return
    }
    // not taking ith character at ith level 
    powerset(res, currentSet, nums, level+1) 
    // taking ith character at ith level 
    powerset(res, append(currentSet, nums[level]), nums, level+1)
}

func subsets(nums []int) [][]int {
    res := make([][]int, 0, len(nums))
    entry := make([]int, 0, len(nums))
    
    backtrack(0, nums, &entry, &res)
    return res
}

func backtrack(start int, nums []int, entry *[]int, res *[][]int) {
    cpy := make([]int, len(*entry))
    copy(cpy, *entry)
    *res = append(*res, cpy)
    
    for i := start; i < len(nums); i++ {
        *entry = append(*entry, nums[i])
        backtrack(i+1, nums, entry, res)
        *entry = (*entry)[:len(*entry)-1]
    }
}