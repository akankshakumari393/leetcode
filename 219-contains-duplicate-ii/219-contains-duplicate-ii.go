func containsNearbyDuplicate(nums []int, k int) bool {
	mapOfElement := make(map[int]int)

	for i1, n := range nums {
		if i2, ok := mapOfElement[n]; ok {
			if i1-i2 <= k {
				return true
			}
		}
		mapOfElement[n] = i1
	}
	return false
}