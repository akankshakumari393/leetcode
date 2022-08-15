func lengthOfLastWord(s string) int {
    wordPos := 0
    for i := len(s)-1; i >=0; i-- {
        if s[i] != 32 {
            wordPos = i
            break
        }
    }
    i := wordPos
    for ; i >= 0; i-- {
        // fmt.Println(s[i])
        if s[i] == 32 {
            break
        }
    }
    return wordPos - i
}