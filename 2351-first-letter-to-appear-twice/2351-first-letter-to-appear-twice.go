func repeatedCharacter(s string) byte {
    //Time Taken O(n)
    freqMap := make(map[byte]int)
    for i, _ := range s {
        freqMap[s[i]]++
        if freqMap[s[i]] == 2 {
            return s[i]
        }
    }
    return ' '
}