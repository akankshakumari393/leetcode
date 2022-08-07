func containsNearbyDuplicate(nums []int, k int) bool {
	mapOfEle := make(map[int]int)

	for i1, n := range nums {
		if i2, ok := mapOfEle[n]; ok {
			if i1-i2 <= k {
				return true
			}
		}
		mapOfEle[n] = i1
	}
	return false
}