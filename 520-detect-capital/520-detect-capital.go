func detectCapitalUse(word string) bool {
    s, _ := regexp.MatchString("^[A-Z]*$|^[a-z]*$|^[A-Z][a-z]*$", word)
	return s
}