func romanToInt(s string) int {
    roman := map[uint8]int{
		73: 1,
		86: 5,
		88: 10,
		76: 50,
		67: 100,
		68: 500,
		77: 1000,
	}
	var a int = 0
	var i uint8 = 0
	for ; i < uint8(len(s))-1; i++ {
		if roman[s[i]] < roman[s[i+1]] {
			a -= roman[s[i]]
		} else {
			a += roman[s[i]]
		}
	}
	a += roman[s[i]]
	return int(a)
}
