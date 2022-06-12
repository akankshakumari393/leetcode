func findDuplicate(nums []int) int {
    slow := nums[0]
    fast := nums[nums[0]]
    for slow != fast {
        slow = nums[slow]
        fast = nums[nums[fast]]
    }
    fast = 0
    for slow != fast {
        slow = nums[slow]
        fast = nums[fast]
    }
    return slow
}