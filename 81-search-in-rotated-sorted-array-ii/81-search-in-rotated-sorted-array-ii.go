func search(nums []int, target int) bool {
    if len(nums) == 0 {
		return false
	}
	left := 0
	right := len(nums) - 1
    //binary seacrh
    for left <= right {
        //find mid
        mid := (left + right) / 2
		if nums[mid] == target {
			return true
		}
        // if left array is sorted
		if nums[left] < nums[mid] {
            // if target lies between left array
			if nums[left] <= target && target <= nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else if nums[left] > nums[mid] { // check for right array sorted
            // if target lies between right array
			if nums[mid] <= target && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
        } else {
            left++ // not sure of the status of error but we are sure that left is equal to mid
        }
	}

	return false
}