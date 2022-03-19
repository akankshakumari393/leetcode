// Time Complexity - O(n)

func merge(nums1 []int, m int, nums2 []int, n int) {
	i := m - 1
	j := n - 1
	for k := m + n - 1; k >= 0; k-- {
        // compare elements backward 
		if (i >= 0 && j >= 0 && nums1[i] > nums2[j]) || (j < 0) {
			nums1[k] = nums1[i]
			i--
		} else {
			nums1[k] = nums2[j]
			j--
		}
	}
}
