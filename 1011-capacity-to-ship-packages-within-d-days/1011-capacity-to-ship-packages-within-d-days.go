// Capacity must be >= 10 as max weight is 10. Cargo must to capable enough to take each weight individually.
// Now we got starting point. What should be our end point?
// sum(weights)? if this can be the capacity of cargo, we can ship everything day 1.

func shipWithinDays(weights []int, days int) int {
    left, right := findMaxSum(weights)   
    for left < right {
        mid := left + (right - left)/2
        
        // calculate the total number of days required to ship with mid capacity of ship
        totalDays := 1
        totalWeight :=0
        for _, weight := range weights {
            totalWeight = totalWeight + weight
            if totalWeight > mid {
                totalDays++
                totalWeight = weight
            }
        }

        if totalDays <= days { //we found the solution but we need to find the minimum
            right = mid
        } else {
            left = mid + 1
        }
    }
    
    return right // right is workable solution    
}

func findMaxSum(weights []int) (int, int) {
    sum := 0
    max := math.MinInt
    for _, weight := range weights {
        sum = sum + weight
        if weight > max {
            max = weight
        }
    }
    return max, sum
}