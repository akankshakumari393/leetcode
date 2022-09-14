func partition(s string) [][]string {
	result := [][]string{}
	current := []string{}
	backtrack(s, 0, &current, &result)
	return result
}

func backtrack(s string, begin int, current *[]string, result *[][]string) {
    if begin >= len(s) {
        *result = append(*result, append([]string{}, *current...))
    }
    // Candidate substring is [begin, end), includes begin, excludes end.
    for end := begin + 1; end <= len(s); end++ {
        // You can save result for each substring in a dp[n][n] to
        // avoid repeating check in the future.
        if isPalindrome(s[begin:end]) {
            *current = append(*current, s[begin:end])
            backtrack(s, end, current, result)
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