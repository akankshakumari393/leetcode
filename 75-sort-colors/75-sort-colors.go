func sortColors(nums []int)  {
    n := len(nums)
    //Time complexity O(2N)
    
    // count0 := 0
    // count1 := 0
    // count2 := 0
    // for _, num := range nums {
    //     if num == 0 {
    //         count0 = count0 + 1 
    //     }else if num == 1 {
    //         count1 = count1 + 1
    //     }else if num == 2 {
    //         count2 = count2 + 1
    //     } 
    // }
    // for i:= 0; i < count0 ; i = i+1 {
    //     nums[i] = 0
    // }
    // for i:= count0; i < count0+count1 ; i = i+1 {
    //     nums[i] = 1
    // }
    // for i:= count0 + count1 ; i < count0 + count1 + count2 ; i = i+1 {
    //     nums[i] = 2
    // }

    // Dutch national flag algorithm
   
    for low, mid, high := 0, 0, n-1 ; mid <= high;  {
        switch nums[mid] {
        case 0 : 
            nums[low], nums[mid] = nums[mid], nums[low]
            mid = mid + 1
            low = low + 1
        case 1 :
            mid = mid + 1
        case 2 :
            nums[high], nums[mid] = nums[mid], nums[high]
            high = high - 1 
        }      
    }
}