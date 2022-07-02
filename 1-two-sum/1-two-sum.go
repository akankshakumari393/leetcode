func twoSum(nums []int, target int) []int {
 // initialize an answer slice
    var ans []int
    numToIdx := make(map[int]int)
    
    for index, num := range nums {
        requiredNum := target - num 
        if value, exists := numToIdx[requiredNum]; exists {
            ans = append(ans, value, index)
        }
        numToIdx[num] = index
    }
    return ans
}