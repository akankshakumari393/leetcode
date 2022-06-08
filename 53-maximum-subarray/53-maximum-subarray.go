func maxSubArray(nums []int) int {
    n := len(nums)
    max := nums[0]
    sum := 0
    // kadane 
    for i :=0; i < n; i++ {
        sum = sum + nums[i]
        if sum > max {
            max = sum
        }
        if sum < 0 {
            sum =0
        }
        //fmt.Println(sum)
    }
return max
}