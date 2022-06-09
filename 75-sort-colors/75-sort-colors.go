func sortColors(nums []int)  {
    //n := len(nums)
    count0 := 0
    count1 := 0
    count2 := 0
    for _, num := range nums {
        if num == 0 {
            count0 = count0 + 1 
        }else if num == 1 {
            count1 = count1 + 1
        }else if num == 2 {
            count2 = count2 + 1
        } 
    }
    // fmt.Println(count0)
    // fmt.Println(count1)
    // fmt.Println(count2)
    for i:= 0; i < count0 ; i = i+1 {
        nums[i] = 0
    }
//        fmt.Println(nums)
    for i:= count0; i < count0+count1 ; i = i+1 {
        nums[i] = 1
    }
// fmt.Println(nums)
    for i:= count0 + count1 ; i < count0 + count1 + count2 ; i = i+1 {
        nums[i] = 2
    }
// fmt.Println(nums)
    
}