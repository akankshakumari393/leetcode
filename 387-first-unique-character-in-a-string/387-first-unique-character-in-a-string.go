//Time Taken O(n)
func firstUniqChar(s string) int {
    freqMap := make(map[byte]int)
    for i, _ := range s {
        freqMap[s[i]]++
    }
    for i, _ := range s {
        if freqMap[s[i]] == 1 {
            return i
        }
    }
    return -1
}