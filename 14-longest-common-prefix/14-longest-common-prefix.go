func longestCommonPrefix(strs []string) string {
	result := strings.Builder{}
	if len(strs) == 0 {
		return result.String()
	}

	for i := 0; i < len(strs[0]); i++ {
		c := strs[0][i]
		same := true
		for _, s := range strs {
            if i >= len(s) || s[i] != c {
				same = false
				break
			}
		}

		if !same {
			break
		}

		result.WriteByte(c)
	}

	return result.String()
}