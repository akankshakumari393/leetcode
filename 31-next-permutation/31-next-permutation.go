func nextPermutation(nums []int)  {
    n := len(nums)
    localMaximaBreak := -1
    for i:= n-2; i>=0; i-- {
        if nums[i] < nums[i+1]{
            localMaximaBreak = i
            for i:= n-1; i>=localMaximaBreak; i-- {
                if nums[i] > nums[localMaximaBreak] {
                    nums[i], nums[localMaximaBreak] = nums[localMaximaBreak], nums[i]
                    goto out
                }
            }
        }
    }
    
    out:
    for i, j :=(localMaximaBreak+1), n-1; i < j; i, j = i+1, j-1 {
        nums[i], nums[j] = nums[j], nums[i]
    }
    
    
    
    
}