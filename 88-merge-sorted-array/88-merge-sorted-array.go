func merge(nums1 []int, m int, nums2 []int, n int)  {
    if (n == 0) {
        return
    }
    
    // check for the first element in 2nd sorted array if greater then swap and then we need to sort the second array because after swaaping that would be unsorted 
    for i := 0; i<m; i++ {
        if nums2[0] < nums1[i] {
            nums1[i], nums2[0] = nums2[0], nums1[i]
            sort(nums2)
        }
    }   
    // copy element from nums2 into nums1 from mth position
    copy(nums1[m:], nums2)
}


func sort(arr []int) {
    // sort the array, Insertion sort trace backward
    for i := 1; i < len(arr); i++ {
        if arr[i-1] > arr[i]  {
            arr[i], arr[i-1] = arr[i-1], arr[i]
        } 
    }
}