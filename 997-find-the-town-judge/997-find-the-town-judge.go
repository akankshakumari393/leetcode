//If any person is truested exactly n-1 times and that doesn't trust anyone then that person is the Judge
func findJudge(n int, trust [][]int) int {
	if n == 1 {
		return 1
	}
	trusteeH := map[int]int{}
	trustSrcH := map[int]int{}

    for _, t := range trust {
		trustSrcH[t[0]]++
		trusteeH[t[1]]++
	}
	
	nM1 := n - 1
	for k := range trusteeH {
		if trusteeH[k] == nM1 && trustSrcH[k] == 0 {
			return k
		}
	}
	return -1    
}

