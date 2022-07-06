import "sort"
func threeSum(nums []int) [][]int {
    var ans [][]int
    // sort the array
    sort.Ints(nums)
    fmt.Println(nums)
    for i:=0; i < len(nums)-2; i++ {
        if (i == 0) || (i>0 && nums[i] != nums[i-1]){
            // Use Two pointer approach
            low, high := i+1, len(nums)-1
            sum := 0 - nums[i]
            for ;low < high; {
                
                if nums[high] + nums[low] == sum {
                    // append the values to ans array
                    ans = append(ans, []int{nums[i], nums[low], nums[high]} )
                    // If next element is same increment the low to avoid processing duplicate element
                    for ;low < high && nums[low] == nums[low + 1];{
                            low++
                    }     
                    // If next element is same decrement the high to avoid processing duplicate element
                    for ;low < high && nums[high] == nums[high-1]; {
                            high--
                    }
                    low ++
                    high --
                } else if nums[low] + nums[high] > sum {
                    high--
                } else {
                    low++
                }
            }             
        }
    }
    return ans
}