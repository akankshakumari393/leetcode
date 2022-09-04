func mySqrt(x int) int {
    left := 0
    right := x
    
    for left <= right {
        mid := left + (right - left)/2
        fmt.Println(left, mid, right)
        val := mid * mid
        nextVal := (mid+1) * (mid+1)
        if x >= val && x < nextVal {
            return mid
        }
        if val > x {
            right = mid - 1
        } else {
            left = mid + 1
        }
    }
    
    return -1
}