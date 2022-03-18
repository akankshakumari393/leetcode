func containsDuplicate(nums []int) bool {
    appearance := make(map[int]bool)

	for _, num := range nums {
		if _, ok := appearance[num]; ok {
			return true
		}
		appearance[num] = true
	}
	return false
}
