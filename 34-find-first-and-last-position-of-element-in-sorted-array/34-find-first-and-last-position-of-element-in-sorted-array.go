func searchRange(nums []int, target int) []int {
    res := []int{-1,-1}
    // find last index
    for left, right := 0, len(nums) - 1; left <= right ;{
        mid := left + (right - left) / 2
        if nums[mid] == target {
            res[1] = mid
        }
        if nums[mid] > target {
            right = mid - 1
        } else {
            left = mid + 1
        }
    }
    // find first index
    for left, right := 0, res[1]; left <= right; {
        mid := left + (right - left) / 2
        if nums[mid] == target{
            res[0] = mid
        }
        if nums[mid] < target {
            left = mid + 1
        } else {
            right = mid - 1
        }
    }
    return res
}