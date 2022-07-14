//If any person is truested exactly n-1 times and that doesn't trust anyone then that person is the Judge
func findJudge(n int, trust [][]int) int {
	if n == 1 {
		return 1
	}
    trustPerson := make([]int, n) 

    for _, t := range trust {
		trustPerson[t[0]-1]--
		trustPerson[t[1]-1]++
	}
	
	for k := range trustPerson {
		if trustPerson[k] == n-1 {
			return k+1
		}
	}
	return -1
}

