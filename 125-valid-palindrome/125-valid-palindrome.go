func isPalindrome(s string) bool {
    n := len(s)
	if n <= 0 {
		return true
	}

	i := 0
	j := n - 1

	for j > i {
		// check character is alphabet
		if (s[i] < 'a' || s[i] > 'z') && (s[i] < 'A' || s[i] > 'Z') && (s[i] < '0' || s[i] > '9') {
			i++
			continue
		}
		if (s[j] < 'a' || s[j] > 'z') && (s[j] < 'A' || s[j] > 'Z')  && (s[j] < '0' || s[j] > '9'){
			j--
			continue
		}

		if strings.ToLower(string(s[i])) != strings.ToLower(string(s[j])) {
			return false
		}

		i++
		j--

	}
	return true
}

