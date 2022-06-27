func singleNonDuplicate(nums []int) int {
    n := len(nums)
    low, high:=0, n-2
    for ; low <= high;  {
        mid := (high+low)/2
        //heavily depend on mid element
        // [3,3,7,7,10,11,11]
                        if (mid % 2 == 0) {
                            if (nums[mid] != nums[mid + 1]) {
                    //Checking whether we are in right half

                        high = mid - 1 //Shrinking the right half
                                
                            } else {
                        low = mid + 1 //Shrinking the left half                                
                            }
                } else {
                     // [1,1,2,3,3,4,4,8,8]
                    //Checking whether we are in right half   
                            if (nums[mid] == nums[mid + 1]) {
                        high = mid - 1 //Shrinking the right half                                
                            } else {
                        low = mid + 1 //Shrinking the left half                                
                            }
                }
    }
    return nums[low]
}