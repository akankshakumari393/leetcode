import "math"
func majorityElement(nums []int) []int {
    // Important point - there can be atmost 2 majority element that appears more than n/3 times.
    element1 := math.MinInt
    element2 := math.MinInt
    c1 := 0
    c2 := 0
    for i:=0; i < len(nums); i++ {
       if element1 != math.MinInt && nums[i] == element1 {
            c1++
        } else if element2 != math.MinInt && nums[i] == element2 {
            c2++
        } else if c1 == 0 {
            element1 = nums[i]
            c1++
        } else if c2 == 0 {
            element2 = nums[i]
            c2++
        } else {
            c1--
            c2--            
        }
    }
    // initialize ans array
    var ans []int
    // reset the count
    c1 = 0
    c2 = 0
    for _, ele := range nums {
        if element1 != math.MinInt && element1 == ele {
            c1++
        }
        if element2 != math.MinInt && element2 == ele {
            c2++
        }
    }
    // check if the count is greater than n/3
    if c1 > len(nums)/3 {
        ans = append(ans, element1)
    }
    if c2 > len(nums)/3 {
        ans = append(ans, element2)
    }
    return ans
}