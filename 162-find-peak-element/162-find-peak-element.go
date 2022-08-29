func findPeakElement(nums []int) int {
    low :=0
    high := len(nums) - 1
    n := len(nums)
    for low <= high {
        mid := low + (high - low)/2 
        // first case if mid is the answer
        if ((mid == 0 || nums[mid - 1] <= nums[mid]) && (mid == n - 1 || nums[mid + 1] <= nums[mid])) {
            return mid   
        } else if (mid > 0 && nums[mid - 1] > nums[mid]) {// move the right pointer
            high = mid - 1   
        } else {
            low = mid + 1; // move the left pointer          
        }
    }
    return -1
}