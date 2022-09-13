var dict = map[byte]string{
	'2': "abc",
	'3': "def",
	'4': "ghi",
	'5': "jkl",
	'6': "mno",
	'7': "pqrs",
	'8': "tuv",
	'9': "wxyz",
}

func letterCombinations(digits string) []string {
    var res []string
    //for input ""
   	if digits == "" {
		return res
	}
    backtrack(0, "", digits, &res)
    return res
}


func backtrack(pos int, entry string, candidates string, res *[]string) {
   // fmt.Println("backtrack called with", entry)
    if pos == len(candidates) {
        *res = append(*res, entry)  
        return
    }
    
    strs := dict[candidates[pos]]
    for i := 0; i < len(strs); i++ {
        backtrack(pos+1, entry+string(strs[i]), candidates, res)
    }
}
