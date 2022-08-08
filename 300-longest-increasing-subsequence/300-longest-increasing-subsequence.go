func lengthOfLIS(nums []int) int {
    tempArr := make([]int, len(nums))
    for i, _ := range tempArr {
        tempArr[i] = 1
    }
    max := 0
    for i, num := range nums {
        for j:=0; j<i; j++ {
            if nums[j] < num {
                if tempArr[j] + 1 > tempArr[i] {
                    tempArr[i] = tempArr[j] + 1
                }
            } 
        }
        if tempArr[i] > max {
            max = tempArr[i]
        }
   }
    return max
}