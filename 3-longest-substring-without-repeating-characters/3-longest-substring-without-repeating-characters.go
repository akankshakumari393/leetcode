func lengthOfLongestSubstring(s string) int {
    left := 0
    right := 0
    maxlen :=0
    posMap := make(map[byte]int, len(s))
    for right < len(s) {
        if i , ok := posMap[s[right]]; ok {
            if i >= left && i <= right {
                left = i + 1
            }
        }
        posMap[s[right]] = right
        len := right -left +1
        if len > maxlen {
            maxlen = len
        }
        right++
    }
    return maxlen
}