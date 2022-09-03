//maxDays - max of the bloomDay  // minDays = min of bloomDay
func minDays(bloomDay []int, m int, k int) int {
    if m*k > len(bloomDay) {
        return -1
    }

    left, right := findMinMax(bloomDay)
    
    for left < right {
        mid := left + (right - left)/2
        bouquets := getPossibleBookies(bloomDay, mid, k)
        if bouquets >= m { //we found the solution but we need to find the minimum
            right = mid
        } else {
            left = mid + 1
        }
    }
    
    return right // right is workable solution
}

// This method is to find the number of bouquets that can be formed on a given day.
func getPossibleBookies(bloomDay []int, day int, k int) int {
    bouquets := 0
    flowersCollected := 0
    for _, value := range bloomDay {
        if (value <= day) {
//If the current flower can be taken with in days then increase the flower flowersCollected.
            flowersCollected++
        } else {
//If there is a flower in between that takes more number of days then the given day, then resent the counter.
            flowersCollected = 0
        }
//If the flowersCollected is same as the required flower per bookie, then increase the bouquets count;
        if flowersCollected == k {
            bouquets++
            flowersCollected = 0
        }
    }
    return bouquets
}

func findMinMax(bloomDay []int) (int, int) {
    min := math.MaxInt
    max := math.MinInt
    for _, day := range bloomDay {
        if day < min {
            min = day
        }
        if day > max {
            max = day
        }
    }
    return min, max
}