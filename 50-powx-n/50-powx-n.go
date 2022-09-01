// binary Search
func myPow(x float64, n int) float64 {
    negative := false
    if n < 0 {
        negative = true
        n = 0-n
    }
    nn := n
    ans := float64(1)
    for nn>0 {
        if (nn % 2 == 1) {
            ans = ans * x
            nn--
        } else {
            x = x*x  // double the x
            nn = nn/2 // half the power
        }
    }
    if negative == true {
        ans = 1/ans
    }
    return ans
}