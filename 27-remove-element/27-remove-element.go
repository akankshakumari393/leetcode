// two pointer approach
func removeElement(nums []int, val int) int {
    lower := 0
    
    for i:=0; i< len(nums); i++ {
        if nums[i] != val {
            nums[lower] = nums[i]
            lower++
        }   
    }
    return lower
}