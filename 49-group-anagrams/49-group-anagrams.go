// TimeComplexity - O(nk)
func groupAnagrams(strs []string) [][]string {
	mp := map[[26]int][]string{}
	for _, s := range strs {
		k := [26]int{}
		for i := 0; i < len(s); i++ {
			k[s[i]-'a'] += 1
		}
		mp[k] = append(mp[k], s)
	}
	res := [][]string{}
	for _, v := range mp {
		res = append(res, v)
	}
	return res
}