func smallestDivisor(nums []int, threshold int) int {
    sort.Ints(nums)

    left := 1
    right := nums[len(nums)-1]
    
    for left < right {
        mid := left + (right - left)/2
        
        // calculate time spent on piles
        result := 0
        for _, num := range nums {
            result += (num/mid)  // find the ceiling
            if num % mid != 0 {
                result++
            }
        }

        if result <= threshold { //we found the solution but we need to find the minimum
            right = mid
        } else {
            left = mid + 1
        }
    }
    
    return right // right is workable solution    
}

// TimeComplexity =naive approach -- max(nums)*n
// TimeComplexity = log(max(pile))