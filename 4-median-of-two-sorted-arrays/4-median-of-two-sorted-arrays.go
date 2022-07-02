func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    return usingTwoPointer(nums1, nums2)
}

func usingTwoPointer(nums1 []int, nums2 []int) float64 {
    // initialize two pointer
    p1 := 0
    p2 := 0
    var ans []int
    for i := 0; p1 < len(nums1) && p2 < len(nums2); i++ {
        if nums1[p1] <= nums2[p2] {
            ans = append(ans, nums1[p1])
            p1++
        } else {
            ans = append(ans, nums2[p2])
            p2++
        }  
    }
    if p1 == len(nums1) {
        ans = append(ans, nums2[p2:]...) 
    } else if p2 == len(nums2){
        ans = append(ans, nums1[p1:]...)         
    }
    fmt.Println(ans)
    
    // find median
    totalLength := len(ans)
    var median float64
    if totalLength%2 == 0 {
        median = float64((ans[(totalLength-1)/2]) + (ans[(totalLength-1)/2 + 1]))/2
    } else {
        median = float64(ans[(totalLength-1)/2])
    }
    return median
}