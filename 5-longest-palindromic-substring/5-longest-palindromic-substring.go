func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}
	l := len(s)
	var maxStart, maxLen int
	for i := 0; i < l; {
		left, right := i, i
		for right < l-1 && s[right] == s[right+1] {
			right++
		}
		i = right + 1
		for left > 0 && right < l-1 && s[left-1] == s[right+1] {
			left--
			right++
		}
		if maxLen < right-left {
			maxStart = left
			maxLen = right - left
		}
	}
	return s[maxStart : maxStart+maxLen+1]
}   
