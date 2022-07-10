import "math"
// At any given position we need to find the minimum of previous two step
// TimeComplexity - O(n)
// Space complexity - O(1)
func minCostClimbingStairs(cost []int) int {
    for i := 2; i < len(cost); i++ {
        cost[i] += int(math.Min(float64(cost[i-1]), float64(cost[i-2])));
    }
    return int(math.Min(float64(cost[len(cost)-1]), float64(cost[len(cost)-2])));
}