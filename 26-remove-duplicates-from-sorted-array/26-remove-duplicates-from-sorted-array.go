// Two pointer Approach
func removeDuplicates(nums []int) int {
    lower := 1
    for i:=1; i< len(nums);i++ {
        if nums[i] != nums[i-1] {
            // new number found
            nums[lower] = nums[i]
            lower++
        }
    }
    return lower
}