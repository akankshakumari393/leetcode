func findJudge(n int, trust [][]int) int {
	if n == 1 {
		return 1
	}
    trustPerson := make([]int, n) 

    for _, t := range trust {
        // the person trust someone
		trustPerson[t[0]-1]--
        // the person is trusted by someone
		trustPerson[t[1]-1]++
	}
	
	for k := range trustPerson {
		if trustPerson[k] == n-1 {
			return k+1
		}
	}
	return -1
}

