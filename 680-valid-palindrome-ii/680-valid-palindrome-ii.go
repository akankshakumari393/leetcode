func validPalindrome(s string) bool {
    l := 0
    r := len(s)-1
    
    for l < r {
        if s[l] != s[r] {
            return isPalindrome(s, l+1, r) || isPalindrome(s, l, r-1)
        } else {
            l++
            r--
        }
    }
    return true
}

func isPalindrome(s string, l, r int) bool {
    for l < r {
        if s[l] != s[r] { return false }
        l++
        r--
    }
    return true
}