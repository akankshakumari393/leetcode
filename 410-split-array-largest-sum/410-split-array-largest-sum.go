func findMaxSum(weights []int) (int, int) {
    sum := 0
    max := 0
    for _, weight := range weights {
        sum = sum + weight
        if weight > max {
            max = weight
        }
    }
    return max, sum
}

func splitArray(nums []int , m int )  (int) {
    // if len(A) < B {
    //     return -1
    // }
    left, right := findMaxSum(nums)
    
    for left < right {
        mid := left + (right - left)/2
        
        // calculate the total number of subarray required to if the sum is mid for each subarray
        totalSubarray := 1
        totalSum :=0
        for _, num := range nums {
            totalSum = totalSum + num
            if totalSum > mid {
                totalSubarray++
                totalSum = num
            }
        }

        if totalSubarray <= m { //we found the solution but we need to find the minimum
            right = mid
        } else {
            left = mid + 1
        }
    }
    
    return right // right is workable solution    
}
