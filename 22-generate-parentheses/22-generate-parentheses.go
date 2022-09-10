func generateParenthesis(n int) []string {
    res := []string{}
    // number of paranthesis left to be printed
	backtrack(&res, "", n, n)
	return res
}

func backtrack(res *[]string, path string, open int, cl int) {
		if open == 0 && cl == 0 {
			*res = append(*res, path)
			return
		}
    // open paranthesis left to be added
		if open > 0 {
			backtrack(res, path+"(", open-1, cl)
		}
    // done with open paranthesis
		if cl > open {
            backtrack(res, path+")", open, cl-1)
		}    
}