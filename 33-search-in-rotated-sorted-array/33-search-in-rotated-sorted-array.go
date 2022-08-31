func search(nums []int, target int) int {
    if len(nums) == 0 {
		return -1
	}
	left := 0
	right := len(nums) - 1
    //binary seacrh
    for left <= right {
        //find mid
        mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		}
        // if left array is sorted
		if nums[left] <= nums[mid] {
            // if target lies between left array
			if nums[left] <= target && target <= nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else { // else right array must be sorted
            // if target lies between right array
			if nums[mid] <= target && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}

	return -1
}