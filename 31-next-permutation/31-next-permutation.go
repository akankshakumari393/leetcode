func nextPermutation(nums []int)  {
    n := len(nums)
    
    localMaximaBreak := -1
    // O(n)
    for i:= n-2; i>=0; i-- {
        if nums[i] < nums[i+1]{
            localMaximaBreak = i
            break
        }
    }
    // o(n)
    if localMaximaBreak != -1 {
        for i:= n-1; i>=localMaximaBreak; i-- {
            if nums[i] > nums[localMaximaBreak] {
                nums[i], nums[localMaximaBreak] = nums[localMaximaBreak], nums[i]
                break
            }
        }
    }
    // this is going to take O(n)
    for i, j :=(localMaximaBreak+1), n-1; i < j; i, j = i+1, j-1 {
        nums[i], nums[j] = nums[j], nums[i]
    }
    
    
    //Time complexity o(1)
    
}