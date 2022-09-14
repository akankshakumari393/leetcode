func partition(s string) [][]string {
	result := [][]string{}
	current := []string{}
	backtrack(0, s, &current, &result)
	return result
}

func backtrack(pos int, s string, current *[]string, result *[][]string) {
    if pos >= len(s) {
        *result = append(*result, append([]string{}, *current...))
    }
    // Candidate substring is [begin, end), includes begin, excludes end.
    for end := pos ; end < len(s); end++ {
        // You can save result for each substring in a dp[n][n] to
        // avoid repeating check in the future.
        if isPalindrome(s[pos:end+1]) {
            *current = append(*current, s[pos:end+1])
            backtrack(end+1, s, current, result)
            *current = (*current)[:len(*current)-1]
        }
    }
    
}


func isPalindrome(str string) bool {
    for i, j := 0, len(str)-1; i < j; i, j = i+1, j-1 {
        if str[i] != str[j] {
            return false
        }
    }
    return true    
}