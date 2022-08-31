func findMin(nums []int) int {
    if len(nums) == 0 {
		return -1
	}
	left := 0
	right := len(nums) - 1
    min := math.MaxInt
    //binary seacrh
    for left <= right {
        //find mid
        mid := (left + right) / 2
        if nums[left] <= nums[mid] && nums[mid] <= nums[right] {
            min = Min(min, nums[left])
		}        
        
        // if left array is sorted 
        if nums[left] <= nums[mid] {
			min = Min(min, nums[left])
            left = mid + 1                
		} else if nums[left] >= nums[mid] { // check for right array sorted
			min = Min(min, nums[mid])
            right = mid - 1
            
	    }
    }

	return min
}

func Min(x, y int) int {
    if x > y {
        return y
    }
    return x
}